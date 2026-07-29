package client

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testCanaryAPIKey = "canary-secret-api-key-do-not-log-12345"

func captureStderr(t *testing.T, fn func()) string {
	t.Helper()

	old := os.Stderr
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stderr = w
	done := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		_ = r.Close()
		done <- buf.String()
	}()
	t.Cleanup(func() {
		os.Stderr = old
		_ = w.Close()
		_ = r.Close()
	})

	fn()
	require.NoError(t, w.Close())
	os.Stderr = old
	return <-done
}

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{}`))
	}))
}

func newHeaderCapturingServer(t *testing.T, captured *string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		*captured = req.Header.Get("X-SpectroCloud-Client")
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{}`))
	}))
}

func assertDebugLogsExcludeCanary(t *testing.T, output string) {
	t.Helper()
	assert.NotContains(t, strings.ToLower(output), strings.ToLower(testCanaryAPIKey),
		"API key must not appear in transport debug logs")
}

func TestWithTransportDebugAtNewRedactsAPIKey(t *testing.T) {
	server := newTestServer(t)
	defer server.Close()

	output := captureStderr(t, func() {
		c := New(
			WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
			WithAPIKey(testCanaryAPIKey),
			WithTransportDebug(),
			WithSchemes([]string{"http"}),
		)
		_, err := c.Client.V1ProjectsMetadata(nil)
		require.NoError(t, err)
	})

	assertDebugLogsExcludeCanary(t, output)
}

func TestWithClientHeaderSetsHeaderOnRequests(t *testing.T) {
	var gotHeader string
	server := newHeaderCapturingServer(t, &gotHeader)
	defer server.Close()

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithSchemes([]string{"http"}),
		WithClientHeader("terraform-v1.2.3"),
	)
	_, err := c.Client.V1ProjectsMetadata(nil)
	require.NoError(t, err)

	assert.Equal(t, "terraform-v1.2.3", gotHeader)
}

func TestWithoutClientHeaderOmitsHeaderOnRequests(t *testing.T) {
	var gotHeader string
	server := newHeaderCapturingServer(t, &gotHeader)
	defer server.Close()

	c := New(
		WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
		WithSchemes([]string{"http"}),
	)
	_, err := c.Client.V1ProjectsMetadata(nil)
	require.NoError(t, err)

	assert.Empty(t, gotHeader)
}

func TestWithTransportDebugAfterNewRedactsAPIKey(t *testing.T) {
	server := newTestServer(t)
	defer server.Close()

	output := captureStderr(t, func() {
		c := New(
			WithPaletteURI(strings.TrimPrefix(server.URL, "http://")),
			WithAPIKey(testCanaryAPIKey),
			WithSchemes([]string{"http"}),
		)
		WithTransportDebug()(c)
		_, err := c.Client.V1ProjectsMetadata(nil)
		require.NoError(t, err)
	})

	assertDebugLogsExcludeCanary(t, output)
}
