package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

type request struct {
	method string
	target string
}

// testRouter is shaped like the API server's: named routes, one carrying a path
// variable, and the middleware attached the same way cmd/server does it.
func testRouter(tb testing.TB) (*mux.Router, *HTTPMetrics) {
	tb.Helper()

	m := NewHTTPMetrics(prometheus.NewRegistry())

	status := func(code int) http.HandlerFunc {
		return func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(code)
		}
	}

	router := mux.NewRouter()
	router.Use(m.Middleware)
	router.NotFoundHandler = m.Middleware(http.NotFoundHandler())

	router.Name("Search.v2").Methods("GET").Path("/v2/search").
		HandlerFunc(status(http.StatusOK))
	router.Name("ingest-file").Methods("POST").Path("/v2/ingest/{fileType}").
		HandlerFunc(status(http.StatusOK))
	router.Name("boom").Methods("GET").Path("/boom").
		HandlerFunc(status(http.StatusInternalServerError))

	return router, m
}

// observations returns how many requests were recorded against the given labels.
func observations(tb testing.TB, m *HTTPMetrics, method, route, code string) uint64 {
	tb.Helper()

	observer, err := m.requestDuration.GetMetricWithLabelValues(method, route, code)
	require.NoError(tb, err)

	metric, ok := observer.(prometheus.Metric)
	require.True(tb, ok)

	var pb dto.Metric
	require.NoError(tb, metric.Write(&pb))
	require.NotNil(tb, pb.Histogram)

	return pb.Histogram.GetSampleCount()
}

func TestMiddleware(t *testing.T) {
	cases := []struct {
		name string

		requests []request

		expectedMethod string
		expectedRoute  string
		expectedCode   string
		expectedCount  uint64
	}{
		{
			name:           "records method, route and code",
			requests:       []request{{"GET", "/v2/search?q=jane"}},
			expectedMethod: "GET",
			expectedRoute:  "/v2/search",
			expectedCode:   "200",
			expectedCount:  1,
		},
		{
			// The query string carries the search term, so it must not reach a label.
			name: "query strings share one series",
			requests: []request{
				{"GET", "/v2/search?q=alice"},
				{"GET", "/v2/search?q=bob"},
			},
			expectedMethod: "GET",
			expectedRoute:  "/v2/search",
			expectedCode:   "200",
			expectedCount:  2,
		},
		{
			name: "path variables collapse to the template",
			requests: []request{
				{"POST", "/v2/ingest/csv"},
				{"POST", "/v2/ingest/json"},
				{"POST", "/v2/ingest/anything-else"},
			},
			expectedMethod: "POST",
			expectedRoute:  "/v2/ingest/{fileType}",
			expectedCode:   "200",
			expectedCount:  3,
		},
		{
			name:           "error status is its own series",
			requests:       []request{{"GET", "/boom"}},
			expectedMethod: "GET",
			expectedRoute:  "/boom",
			expectedCode:   "500",
			expectedCount:  1,
		},
		{
			name:           "unmatched requests are recorded",
			requests:       []request{{"GET", "/no/such/path"}},
			expectedMethod: "GET",
			expectedRoute:  unmatchedRoute,
			expectedCode:   "404",
			expectedCount:  1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			router, m := testRouter(t)

			for _, req := range tc.requests {
				w := httptest.NewRecorder()
				router.ServeHTTP(w, httptest.NewRequest(req.method, req.target, nil))
			}

			got := observations(t, m, tc.expectedMethod, tc.expectedRoute, tc.expectedCode)
			require.Equal(t, tc.expectedCount, got)
		})
	}
}

func TestMethodLabel(t *testing.T) {
	cases := []struct {
		name string

		method   string
		expected string
	}{
		{name: "get", method: http.MethodGet, expected: "GET"},
		{name: "post", method: http.MethodPost, expected: "POST"},
		{name: "delete", method: http.MethodDelete, expected: "DELETE"},

		// Anything a caller invents collapses, so it cannot mint time series.
		{name: "invented", method: "BREW", expected: "other"},
		{name: "empty", method: "", expected: "other"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, methodLabel(tc.method))
		})
	}
}

// Handlers asserting http.Flusher directly must still find it — embedding
// http.ResponseWriter in a wrapper silently strips it.
func TestMiddleware_PreservesFlusher(t *testing.T) {
	m := NewHTTPMetrics(prometheus.NewRegistry())

	var sawFlusher bool
	handler := m.Middleware(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, sawFlusher = w.(http.Flusher)
		w.WriteHeader(http.StatusOK)
	}))

	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)

	resp, err := http.Get(server.URL)
	require.NoError(t, err)
	t.Cleanup(func() { resp.Body.Close() })

	require.True(t, sawFlusher, "handler lost http.Flusher")
}
