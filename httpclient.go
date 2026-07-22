package hotels

import (
	"log/slog"
	"net/http"
	"time"
)

type httpclient struct {
	next http.RoundTripper
}

func (t *httpclient) RoundTrip(req *http.Request) (*http.Response, error) {
	slog.InfoContext(req.Context(), "outgoing request",
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()))

	start := time.Now()
	resp, err := t.next.RoundTrip(req)

	slog.InfoContext(req.Context(), "outgoing request completed",
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		slog.Duration("duration", time.Since(start)),
		slog.Any("error", err))

	return resp, err
}

func NewHTTPClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout:   timeout,
		Transport: &httpclient{next: http.DefaultTransport},
	}
}
