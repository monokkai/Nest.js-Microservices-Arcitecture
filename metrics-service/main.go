package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests",
		}, []string{"path", "method"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func main() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestDuration)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(50 * time.Millisecond)

		httpRequestsTotal.WithLabelValues(r.Method, "/health").Inc()
		httpRequestDuration.WithLabelValues(r.Method, "/health").
			Observe(time.Since(start).Seconds())

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("metrics-service running on :9091")
	log.Fatal(http.ListenAndServe(":9091", nil))
}
