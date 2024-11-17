package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	s := http.NewServeMux()

	scope := "local"
	if customScope := os.Getenv("SCOPE"); customScope != "" {
		scope = customScope
	}

	counter := promauto.NewCounter(prometheus.CounterOpts{
		Name: "simple_ping_request_done",
		ConstLabels: map[string]string{
			"scope": scope,
		},
	})

	s.Handle("/metrics", promhttp.Handler())
	s.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		counter.Inc()
		w.Write([]byte("pong"))
	})

	port := "8080"
	if customPort := os.Getenv("PORT"); customPort != "" {
		port = customPort
	}

	locationCertificate := os.Getenv("CERT_FILE_LOCATION")
	locationKey := os.Getenv("KEY_FILE_LOCATION")

	fmt.Printf("Running in port %s!\n", port)

	if err := http.ListenAndServeTLS(":"+port, locationCertificate, locationKey, s); err != nil {
		fmt.Println(err)
	}
}
