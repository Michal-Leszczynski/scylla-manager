// Copyright (C) 2017 ScyllaDB

package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/hailocab/go-hostpool"
	"github.com/pkg/errors"
	"github.com/scylladb/go-log"
	httputilx "github.com/scylladb/mermaid/internal/httputil"
	"github.com/scylladb/mermaid/internal/retryablehttp"
	"github.com/scylladb/mermaid/internal/timeutc"
)

// FixContentType adjusts Scylla REST API response so that it can be consumed
// by Open API.
func FixContentType(next http.RoundTripper) http.RoundTripper {
	return httputilx.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
		defer func() {
			if resp != nil {
				// Force JSON, Scylla returns "text/plain" that misleads the
				// unmarshaller and breaks processing.
				resp.Header.Set("Content-Type", "application/json")
			}
		}()
		return next.RoundTrip(req)
	})
}

// Retry retries request if needed.
func Retry(next http.RoundTripper, poolSize int, logger log.Logger) http.RoundTripper {
	// Retry policy while using a specified host.
	hostRetry := retryablehttp.NewTransport(next, logger)

	// Retry policy while using host pool, no backoff wait time, switch host as
	// fast as possible.
	poolRetry := retryablehttp.NewTransport(next, logger)
	poolRetry.Backoff = func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration { return 0 }
	poolRetry.RetryMax = poolSize

	return httputilx.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
		if _, ok := req.Context().Value(ctxNoRetry).(bool); ok {
			return next.RoundTrip(req)
		}

		if _, ok := req.Context().Value(ctxHost).(string); ok {
			return hostRetry.RoundTrip(req)
		}

		return poolRetry.RoundTrip(req)
	})
}

// HostPool sets request host from a pool.
func HostPool(next http.RoundTripper, pool hostpool.HostPool, port string) http.RoundTripper {
	return httputilx.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
		ctx := req.Context()

		var (
			h   string
			hpr hostpool.HostPoolResponse
		)

		// Get host from context
		h, ok := ctx.Value(ctxHost).(string)

		// Get host from pool
		if !ok {
			hpr = pool.Get()
			h = hpr.Host()
		}

		// Clone request
		r := cloneRequest(req)

		// Set host and port
		hp := h + ":" + port
		r.Host = hp
		r.URL.Host = hp

		// RoundTrip shall not modify requests, here we modify it to fix error
		// messages see https://github.com/scylladb/mermaid/issues/266.
		// This is legit because we own the whole process. The modified request
		// is not being sent.
		req.Host = h
		req.URL.Host = h

		resp, err := next.RoundTrip(r)

		// Mark response
		if hpr != nil {
			hpr.Mark(err)
		}

		return resp, err
	})
}

// Logger logs requests and responses.
func Logger(next http.RoundTripper, logger log.Logger) http.RoundTripper {
	return httputilx.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
		start := timeutc.Now()
		defer func() {
			if resp != nil {
				f := []interface{}{
					"host", req.Host,
					"method", req.Method,
					"uri", req.URL.RequestURI(),
					"status", resp.StatusCode,
					"bytes", resp.ContentLength,
					"duration", fmt.Sprintf("%dms", timeutc.Since(start)/1000000),
				}

				// Dump body of failed requests
				if resp.StatusCode >= 400 {
					if b, err := httputil.DumpResponse(resp, true); err != nil {
						f = append(f, "dump", errors.Wrap(err, "failed to dump request"))
					} else {
						f = append(f, "dump", string(b))
					}
					logger.Info(req.Context(), "HTTP", f...)
				} else {
					logger.Debug(req.Context(), "HTTP", f...)
				}
			}
		}()
		return next.RoundTrip(req)
	})
}

// body defers context cancellation until response body is closed.
type body struct {
	io.ReadCloser
	cancel context.CancelFunc
}

func (b body) Close() error {
	defer b.cancel()
	return b.ReadCloser.Close()
}

// Timeout sets request context timeout for individual requests.
func Timeout(next http.RoundTripper, timeout time.Duration) http.RoundTripper {
	return httputilx.RoundTripperFunc(func(req *http.Request) (resp *http.Response, err error) {
		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer func() {
			if resp != nil {
				resp.Body = body{resp.Body, cancel}
			}
		}()
		return next.RoundTrip(req.WithContext(ctx))
	})
}

// cloneRequest creates a shallow copy of the request along with a deep copy
// of the Headers and URL.
func cloneRequest(req *http.Request) *http.Request {
	r := new(http.Request)

	// Shallow clone
	*r = *req

	// Deep copy headers
	r.Header = cloneHeader(req.Header)

	// Deep copy URL
	r.URL = new(url.URL)
	*r.URL = *req.URL

	return r
}

// cloneHeader creates a deep copy of an http.Header.
func cloneHeader(in http.Header) http.Header {
	out := make(http.Header, len(in))
	for key, values := range in {
		newValues := make([]string, len(values))
		copy(newValues, values)
		out[key] = newValues
	}
	return out
}