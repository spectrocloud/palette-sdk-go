package transport

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"

	"github.com/go-openapi/runtime"
	openapiclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const canaryAPIKey = "canary-secret-api-key-do-not-log-12345"

type captureLogger struct {
	mu    sync.Mutex
	lines []string
}

func (l *captureLogger) Printf(format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lines = append(l.lines, fmt.Sprintf(format, args...))
}

func (l *captureLogger) Debugf(format string, args ...interface{}) {
	l.Printf(format, args...)
}

func (l *captureLogger) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return strings.Join(l.lines, "\n")
}

func TestDebugDumpRedactsAPIKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set(runtime.HeaderContentType, runtime.JSONMime)
		rw.Header().Set("Set-Cookie", "session=super-secret-cookie")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"status":"ok"}`))
	}))
	defer server.Close()

	log := &captureLogger{}
	hu, err := url.Parse(server.URL)
	require.NoError(t, err)

	rt := New(hu.Host, "/", []string{"http"})
	rt.SetDebug(true)
	rt.SetLogger(log)
	rt.DefaultAuthentication = openapiclient.APIKeyAuth("ApiKey", "header", canaryAPIKey)
	rt.AddSensitiveValue(canaryAPIKey)

	_, err = rt.Submit(&runtime.ClientOperation{
		Method:      http.MethodGet,
		PathPattern: "/",
		Params: runtime.ClientRequestWriterFunc(func(_ runtime.ClientRequest, _ strfmt.Registry) error {
			return nil
		}),
		Reader: runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
			if response.Code() != http.StatusOK {
				return nil, assert.AnError
			}
			var res map[string]string
			if err := consumer.Consume(response.Body(), &res); err != nil {
				return nil, err
			}
			return res, nil
		}),
	})
	require.NoError(t, err)

	output := log.String()
	assert.NotContains(t, output, canaryAPIKey, "API key must not appear in debug logs")
	assert.NotContains(t, output, "super-secret-cookie", "Set-Cookie must not appear in debug logs")
	assert.Contains(t, output, "GET /")
}

func TestDebugDumpRedactsSensitiveBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set(runtime.HeaderContentType, runtime.JSONMime)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"token":"ok"}`))
	}))
	defer server.Close()

	const bodySecret = "body-secret-password-value"
	log := &captureLogger{}
	hu, err := url.Parse(server.URL)
	require.NoError(t, err)

	rt := New(hu.Host, "/", []string{"http"})
	rt.SetDebug(true)
	rt.SetLogger(log)
	rt.AddSensitiveValue(bodySecret)

	_, err = rt.Submit(&runtime.ClientOperation{
		Method:      http.MethodPost,
		PathPattern: "/",
		Params: runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
			return req.SetBodyParam(map[string]string{"password": bodySecret})
		}),
		Reader: runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
			if response.Code() != http.StatusOK {
				return nil, assert.AnError
			}
			var res map[string]string
			if err := consumer.Consume(response.Body(), &res); err != nil {
				return nil, err
			}
			return res, nil
		}),
	})
	require.NoError(t, err)

	output := log.String()
	assert.NotContains(t, output, bodySecret, "sensitive body values must be redacted in debug logs")
	assert.Contains(t, output, "[REDACTED]")
}

func TestSanitizeAuthHeaders(t *testing.T) {
	h := http.Header{
		"ApiKey":        []string{canaryAPIKey},
		"Authorization": []string{"Bearer jwt-token"},
		"Content-Type":  []string{"application/json"},
	}
	sanitizeAuthHeaders(h)
	assert.Empty(t, h.Get("ApiKey"))
	assert.Empty(t, h.Get("Authorization"))
	assert.Equal(t, "application/json", h.Get("Content-Type"))
}

func TestLogResponseDebugPreservesBody(t *testing.T) {
	log := &captureLogger{}
	rt := &Runtime{Debug: true, logger: log}
	rt.AddSensitiveValue("response-body-secret")

	body := []byte(`{"secret":"response-body-secret"}`)
	res := &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Set-Cookie": []string{"sid=abc"}},
		Body:       io.NopCloser(bytes.NewReader(body)),
	}
	require.NoError(t, rt.logResponseDebug(res))

	restored, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	assert.Equal(t, body, restored)

	output := log.String()
	assert.NotContains(t, output, "response-body-secret")
	assert.NotContains(t, output, "sid=abc")
}
