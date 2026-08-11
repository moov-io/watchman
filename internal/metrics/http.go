// Package metrics publishes Prometheus metrics describing the HTTP traffic that
// watchman serves. Collectors registered here are exposed by the admin server on
// /metrics, alongside the Go runtime and process collectors.
package metrics

import (
	"net/http"
	"strconv"

	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// DefaultBuckets covers the range a search spends: the scoring loop runs over the
// candidate set in memory, so the interesting region is milliseconds to a few seconds.
// These are the Prometheus defaults, which fit that shape.
var DefaultBuckets = []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10}

// HTTPMetrics collects timings for the API server's routes.
type HTTPMetrics struct {
	requestDuration *prometheus.HistogramVec
}

// NewHTTPMetrics registers the HTTP collectors against reg. Pass
// prometheus.DefaultRegisterer to have them served by the admin server.
func NewHTTPMetrics(reg prometheus.Registerer) *HTTPMetrics {
	return &HTTPMetrics{
		requestDuration: promauto.With(reg).NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "watchman",
			Subsystem: "http",
			Name:      "request_duration_seconds",
			Help:      "Duration of HTTP requests served by the API server, in seconds",
			Buckets:   DefaultBuckets,
		}, []string{"method", "route", "code"}),
	}
}

// Middleware times each request and records it against the route that served it.
// It satisfies mux.MiddlewareFunc, so it can be attached with router.Use.
func (m *HTTPMetrics) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// httpsnoop keeps the optional interfaces a plain wrapper would strip.
		result := httpsnoop.CaptureMetrics(next, w, r)

		m.requestDuration.WithLabelValues(
			methodLabel(r.Method),
			routeLabel(r),
			strconv.Itoa(result.Code),
		).Observe(result.Duration.Seconds())
	})
}

// routeLabel reports the registered path template rather than the requested path, so
// that a route carrying variables — /v2/ingest/{fileType} — stays a single time series
// instead of one per value seen.
func routeLabel(r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route == nil {
		return unmatchedRoute
	}
	if tpl, err := route.GetPathTemplate(); err == nil && tpl != "" {
		return tpl
	}
	if name := route.GetName(); name != "" {
		return name
	}
	return unmatchedRoute
}

const unmatchedRoute = "unmatched"

// methodLabel bounds the label to the methods net/http defines. The request method is
// caller-controlled, and an arbitrary token would otherwise mint a time series per value.
func methodLabel(method string) string {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut,
		http.MethodPatch, http.MethodDelete, http.MethodConnect,
		http.MethodOptions, http.MethodTrace:
		return method
	}
	return "other"
}
