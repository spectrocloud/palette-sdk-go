package client

import (
	"crypto/x509"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testScheme = "https"

func newTLSTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	server := httptest.NewTLSServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{}`))
	}))
	t.Cleanup(server.Close)
	return server
}

func poolFor(servers ...*httptest.Server) *x509.CertPool {
	pool := x509.NewCertPool()
	for _, s := range servers {
		pool.AddCert(s.Certificate())
	}
	return pool
}

func hostOf(t *testing.T, server *httptest.Server) string {
	t.Helper()
	return strings.TrimPrefix(server.URL, "https://")
}

// canReach returns true if a client built with the given options can
// successfully complete a request against server, false if the TLS
// handshake (or the request itself) fails.
func canReach(t *testing.T, server *httptest.Server, opts ...func(*V1Client)) bool {
	t.Helper()
	base := []func(*V1Client){
		WithPaletteURI(hostOf(t, server)),
		WithSchemes([]string{testScheme}),
	}
	c := New(append(base, opts...)...)
	_, err := c.Client.V1ProjectsMetadata(nil)
	return err == nil
}

func TestHTTPClient_NilRootCAs_UsesSystemDefault(t *testing.T) {
	server := newTLSTestServer(t)

	// No WithRootCAs option at all — a self-signed test-server cert is not
	// in the system trust store, so the request must fail. This is the
	// SaaS-path regression check: existing callers who never touch this
	// option must see unchanged, strict verification.
	assert.False(t, canReach(t, server), "client with no RootCAs must not trust a self-signed cert")
}

func TestHTTPClient_RootCAs_TrustsOnlySuppliedCA(t *testing.T) {
	server := newTLSTestServer(t)
	pool := poolFor(server)

	assert.True(t, canReach(t, server, WithRootCAs(pool)), "client with the server's CA in RootCAs must trust it")
}

func TestWithRootCAs_ClonesInput(t *testing.T) {
	serverA := newTLSTestServer(t)
	serverB := newTLSTestServer(t)

	pool := poolFor(serverA)
	c := New(
		WithPaletteURI(hostOf(t, serverA)),
		WithSchemes([]string{testScheme}),
		WithRootCAs(pool),
	)

	// Mutate the caller's pool AFTER handing it to WithRootCAs. If the
	// option stored the pointer directly (no defensive clone), this
	// mutation leaks into the client's trust store.
	pool.AddCert(serverB.Certificate())

	_, err := c.Client.V1ProjectsMetadata(nil)
	assert.NoError(t, err, "client must still trust serverA after the external pool was mutated")

	// Directly verify c's internal pool is a distinct object from the
	// caller's pool — the mutation above must not have reached it.
	assert.NotSame(t, pool, c.rootCAs, "WithRootCAs must clone the supplied pool, not store the caller's pointer")
}

func TestClone_PropagatesRootCAs(t *testing.T) {
	server := newTLSTestServer(t)
	pool := poolFor(server)

	c := New(
		WithPaletteURI(hostOf(t, server)),
		WithSchemes([]string{testScheme}),
		WithRootCAs(pool),
	)
	_, err := c.Client.V1ProjectsMetadata(nil)
	require.NoError(t, err, "sanity: original client must trust the server before cloning")

	cloned := c.Clone()
	// Clone() rebuilds paletteURI/scheme from tenant scope defaults, so the
	// meaningful assertion isn't "same host" (Clone always resets to
	// WithScopeTenant with the original paletteURI) — it's that the CLONE
	// actually carries rootCAs at all. Before the fix, Clone()'s option
	// list never included WithRootCAs, so cloned.rootCAs was always nil
	// regardless of h.rootCAs, and the clone would fail to trust anything
	// beyond the system pool even though the original client worked fine.
	require.NotNil(t, cloned.rootCAs, "Clone() must propagate rootCAs — this was the bug: Clone() never called WithRootCAs")
	assert.NotSame(t, c.rootCAs, cloned.rootCAs, "Clone() must clone the pool again, not share the original client's pointer")
}

func TestClone_PlainSaaSClient_NoRootCAs_Unaffected(t *testing.T) {
	// A client that never set WithRootCAs (the plain-SaaS case, unchanged
	// by this patch) must still Clone() cleanly with rootCAs left nil —
	// the `if h.rootCAs != nil` guard in Clone() must not, say, panic on a
	// nil pool or otherwise regress the pre-existing SaaS Clone() path.
	c := New(
		WithPaletteURI("console.spectrocloud.com"),
		WithAPIKey("fake-key-not-used-over-the-wire"),
	)
	if c.rootCAs != nil {
		t.Fatalf("sanity: expected nil rootCAs on a client that never called WithRootCAs, got %v", c.rootCAs)
	}

	cloned := c.Clone()
	assert.Nil(t, cloned.rootCAs, "Clone() of a plain-SaaS client must leave rootCAs nil, not introduce one")
}
