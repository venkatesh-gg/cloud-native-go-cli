package metrics

import (
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

// RequestCount counts HTTP requests by path.
var RequestCount = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "app_requests_total",
        Help: "Total number of HTTP requests",
    },
    []string{"path"},
)

// Register registers all Prometheus metrics.
func Register() {
    prometheus.MustRegister(RequestCount)
}

// WrapHandler instruments a handler to count requests by path.
// It manually increments the counter for each request.
func WrapHandler(path string, h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        RequestCount.WithLabelValues(path).Inc()
        h.ServeHTTP(w, r)
    })
}

// Handler returns the HTTP handler for /metrics.
func Handler() http.Handler {
    return promhttp.Handler()
}

