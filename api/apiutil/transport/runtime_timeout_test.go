package transport

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PLT-2495: an API call whose response stalls on the network must fail within a bounded time instead of blocking
// the caller forever.

// stallServer answers with 200 and part of a JSON body, then sends nothing more until the client goes away. It
// counts the requests it receives.
func stallServer(t *testing.T, hits *atomic.Int32) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		hits.Add(1)
		rw.Header().Set(runtime.HeaderContentType, runtime.JSONMime)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(`{"status":`))
		rw.(http.Flusher).Flush()
		<-req.Context().Done()
	}))
	t.Cleanup(srv.Close)
	return srv
}

func newTimeoutTestRuntime(t *testing.T, srv *httptest.Server) *Runtime {
	t.Helper()
	hu, err := url.Parse(srv.URL)
	require.NoError(t, err)
	rt := New(hu.Host, "/", []string{"http"})
	rt.RetryAttempts = 0
	return rt
}

// timeoutTestOperation mirrors a generated operation built with New...ParamsWithContext: the parameters set no
// timeout of their own, so the runtime's timeouts apply.
func timeoutTestOperation(ctx context.Context, method string) *runtime.ClientOperation {
	return &runtime.ClientOperation{
		Method:             method,
		PathPattern:        "/v1/stall",
		ProducesMediaTypes: []string{runtime.JSONMime},
		ConsumesMediaTypes: []string{runtime.JSONMime},
		Schemes:            []string{"http"},
		Context:            ctx,
		Params: runtime.ClientRequestWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			return r.SetTimeout(0)
		}),
		Reader: runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
			var res map[string]string
			if err := consumer.Consume(response.Body(), &res); err != nil {
				return nil, err
			}
			return res, nil
		}),
	}
}

func submitWithin(t *testing.T, limit time.Duration, rt *Runtime, op *runtime.ClientOperation) (interface{}, error) {
	t.Helper()
	type outcome struct {
		res interface{}
		err error
	}
	done := make(chan outcome, 1)
	go func() {
		res, err := rt.Submit(op)
		done <- outcome{res, err}
	}()
	select {
	case o := <-done:
		return o.res, o.err
	case <-time.After(limit):
		t.Fatalf("Submit still blocked after %s", limit)
		return nil, nil
	}
}

func TestSubmitTimesOutStalledBodyWithoutCallerDeadline(t *testing.T) {
	var hits atomic.Int32
	rt := newTimeoutTestRuntime(t, stallServer(t, &hits))
	rt.RequestTimeout = 300 * time.Millisecond
	rt.ResponseIdleTimeout = -1 // only the request timeout can end the call

	_, err := submitWithin(t, 10*time.Second, rt, timeoutTestOperation(context.Background(), http.MethodGet))
	require.Error(t, err)
}

func TestSubmitStalledBodyFailsAfterResponseIdleTimeout(t *testing.T) {
	var hits atomic.Int32
	rt := newTimeoutTestRuntime(t, stallServer(t, &hits))
	rt.RequestTimeout = time.Minute
	rt.ResponseIdleTimeout = 200 * time.Millisecond

	start := time.Now()
	_, err := submitWithin(t, 10*time.Second, rt, timeoutTestOperation(context.Background(), http.MethodGet))
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrResponseStalled)
	assert.Less(t, time.Since(start), 5*time.Second, "must fail at the idle timeout, well before the request timeout")
}

func TestSubmitSlowServerAndTrickleBodySucceed(t *testing.T) {
	// The server takes longer than the idle timeout to answer, then sends the body in small pieces, each within the
	// idle timeout. Neither may trip it: the timer starts at the headers and restarts on every piece.
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		time.Sleep(800 * time.Millisecond)
		rw.Header().Set(runtime.HeaderContentType, runtime.JSONMime)
		rw.WriteHeader(http.StatusOK)
		for _, part := range []string{`{"sta`, `tus":`, `"ok"`, `}`} {
			_, _ = rw.Write([]byte(part))
			rw.(http.Flusher).Flush()
			time.Sleep(150 * time.Millisecond)
		}
	}))
	t.Cleanup(srv.Close)
	rt := newTimeoutTestRuntime(t, srv)
	rt.RequestTimeout = time.Minute
	rt.ResponseIdleTimeout = 300 * time.Millisecond

	res, err := submitWithin(t, 10*time.Second, rt, timeoutTestOperation(context.Background(), http.MethodGet))
	require.NoError(t, err)
	assert.Equal(t, map[string]string{"status": "ok"}, res)
}

func TestSubmitHonorsCallerDeadline(t *testing.T) {
	var hits atomic.Int32
	rt := newTimeoutTestRuntime(t, stallServer(t, &hits))
	rt.RequestTimeout = time.Hour
	rt.ResponseIdleTimeout = -1
	rt.RetryAttempts = 2

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := submitWithin(t, 10*time.Second, rt, timeoutTestOperation(ctx, http.MethodGet))
	require.Error(t, err)
	assert.Less(t, time.Since(start), 2*time.Second, "the caller's deadline must end the call")
	assert.Equal(t, int32(1), hits.Load(), "a call whose caller gave up must not be retried")
}

func TestSubmitRetriesTimedOutReadsOnly(t *testing.T) {
	tests := []struct {
		method   string
		wantHits int32
	}{
		{method: http.MethodGet, wantHits: 2},
		{method: http.MethodPost, wantHits: 1},
	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			var hits atomic.Int32
			rt := newTimeoutTestRuntime(t, stallServer(t, &hits))
			rt.ResponseIdleTimeout = 200 * time.Millisecond
			rt.RetryAttempts = 1

			_, err := submitWithin(t, 20*time.Second, rt, timeoutTestOperation(context.Background(), tt.method))
			require.Error(t, err)
			assert.ErrorIs(t, err, ErrResponseStalled)
			assert.Equal(t, tt.wantHits, hits.Load())
		})
	}
}

func TestSubmitRequestTimeoutBeforeHeadersIsDeadlineExceeded(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, req *http.Request) {
		<-req.Context().Done() // never answers
	}))
	t.Cleanup(srv.Close)
	rt := newTimeoutTestRuntime(t, srv)
	rt.RequestTimeout = 200 * time.Millisecond

	_, err := submitWithin(t, 10*time.Second, rt, timeoutTestOperation(context.Background(), http.MethodGet))
	require.Error(t, err)
	var tcpErr *TcpError
	require.True(t, errors.As(err, &tcpErr))
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestDefaultTimeouts(t *testing.T) {
	assert.Equal(t, 5*time.Minute, DefaultRequestTimeout)
	assert.Equal(t, 60*time.Second, DefaultResponseIdleTimeout)

	rt := New("example.com", "/", []string{"https"})
	req := &request{}
	assert.Equal(t, DefaultRequestTimeout, rt.getRequestTimeout(req), "zero uses the default")
	assert.Equal(t, DefaultResponseIdleTimeout, rt.getResponseIdleTimeout(), "zero uses the default")

	rt.RequestTimeout = 2 * time.Minute
	rt.ResponseIdleTimeout = 10 * time.Second
	assert.Equal(t, 2*time.Minute, rt.getRequestTimeout(req))
	assert.Equal(t, 10*time.Second, rt.getResponseIdleTimeout())

	req.timeout = 10 * time.Minute
	assert.Equal(t, 10*time.Minute, rt.getRequestTimeout(req), "a longer timeout on the request parameters wins")

	rt.RequestTimeout = -1
	rt.ResponseIdleTimeout = -1
	assert.Equal(t, 10*time.Minute, rt.getRequestTimeout(req), "disabling the default keeps the request's own timeout")
	assert.Equal(t, time.Duration(0), rt.getRequestTimeout(&request{}), "negative disables the limit")
	assert.Equal(t, time.Duration(0), rt.getResponseIdleTimeout(), "negative disables the check")
}

func TestIsIdempotentMethod(t *testing.T) {
	for _, m := range []string{http.MethodGet, http.MethodHead, strings.ToLower(http.MethodGet)} {
		assert.True(t, isIdempotentMethod(m), m)
	}
	for _, m := range []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete} {
		assert.False(t, isIdempotentMethod(m), m)
	}
}
