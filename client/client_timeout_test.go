package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spectrocloud/palette-sdk-go/api/apiutil/transport"
)

// PLT-2495: a client call whose response stalls must fail instead of blocking the caller forever. Before the fix,
// GetClusterWithoutStatus against this server never returned.
func TestStalledResponseBodyFails(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		hits.Add(1)
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"metadata":{"uid":`))
		rw.(http.Flusher).Flush()
		<-req.Context().Done()
	}))
	defer srv.Close()

	c := New(
		WithPaletteURI(strings.TrimPrefix(srv.URL, "http://")),
		WithSchemes([]string{"http"}),
		WithRetries(0),
		WithResponseIdleTimeout(200*time.Millisecond),
	)

	done := make(chan error, 1)
	go func() {
		_, err := c.GetClusterWithoutStatus("cluster-uid")
		done <- err
	}()
	select {
	case err := <-done:
		require.Error(t, err)
		assert.ErrorIs(t, err, transport.ErrResponseStalled)
	case <-time.After(10 * time.Second):
		t.Fatal("GetClusterWithoutStatus still blocked on a stalled response")
	}
	assert.Equal(t, int32(1), hits.Load())
}

func TestTimeoutOptionsReachTransport(t *testing.T) {
	c := New(
		WithPaletteURI("api.example.com"),
		WithRequestTimeout(2*time.Minute),
		WithResponseIdleTimeout(-1),
	)
	rt := c.baseTransport()
	assert.Equal(t, 2*time.Minute, rt.RequestTimeout)
	assert.Equal(t, time.Duration(-1), rt.ResponseIdleTimeout)

	defaults := New(WithPaletteURI("api.example.com")).baseTransport()
	assert.Zero(t, defaults.RequestTimeout, "unset keeps transport.DefaultRequestTimeout")
	assert.Zero(t, defaults.ResponseIdleTimeout, "unset keeps transport.DefaultResponseIdleTimeout")
}

func TestClone_PreservesTimeouts(t *testing.T) {
	orig := New(
		WithPaletteURI("api.example.com"),
		WithAPIKey("k"),
		WithRequestTimeout(7*time.Minute),
		WithResponseIdleTimeout(90*time.Second),
	)
	cloned := orig.Clone()
	assert.Equal(t, 7*time.Minute, cloned.requestTimeout)
	assert.Equal(t, 90*time.Second, cloned.responseIdleTimeout)
}
