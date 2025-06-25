package main

import (
	"log"
	"net/http"

	"github.com/venkateshgogula/cloud-native-go-cli/pkg/config"
	"github.com/venkateshgogula/cloud-native-go-cli/pkg/metrics"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Register Prometheus metrics
	metrics.Register()

	// Health endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Metrics endpoint
	http.Handle("/metrics", metrics.Handler())

	// Example instrumented endpoint
	hello := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Cloud Service"))
	})
	http.Handle("/hello", metrics.WrapHandler("/hello", hello))

	log.Printf("🚀 Server listening on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, nil))
}
