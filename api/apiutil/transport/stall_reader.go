package transport

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

// ErrResponseStalled is returned when a response body stops delivering data for longer than the response idle timeout.
var ErrResponseStalled = errors.New("response stalled: no data received")

// stallReader wraps a response body and cancels the request when no bytes arrive for idle. The timer starts once the
// response headers have been received and restarts on every read that returns data, so a slow but steadily arriving
// body never trips it.
type stallReader struct {
	body    io.ReadCloser
	idle    time.Duration
	timer   *time.Timer
	stalled atomic.Bool
}

func newStallReader(body io.ReadCloser, idle time.Duration, cancel context.CancelCauseFunc) *stallReader {
	s := &stallReader{body: body, idle: idle}
	s.timer = time.AfterFunc(idle, func() {
		s.stalled.Store(true)
		cancel(ErrResponseStalled)
	})
	return s
}

func (s *stallReader) Read(p []byte) (int, error) {
	n, err := s.body.Read(p)
	if n > 0 && !s.stalled.Load() {
		s.timer.Reset(s.idle)
	}
	switch {
	case err == nil:
	case errors.Is(err, io.EOF):
		s.timer.Stop()
	case s.stalled.Load():
		err = fmt.Errorf("%w for %s: %w", ErrResponseStalled, s.idle, err)
	}
	return n, err
}

func (s *stallReader) Close() error {
	s.timer.Stop()
	return s.body.Close()
}
