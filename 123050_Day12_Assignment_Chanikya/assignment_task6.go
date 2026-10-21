package main

import (
	"log"
	"net/http"
	"time"
)

// Task 6: Operational microsecond logging middleware
// Captures time metrics with high accuracy to identify structural application latencies.

func MicrosecondLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Microseconds()
		log.Printf("[HTTP] %s %s took %d microseconds", r.Method, r.URL.Path, duration)
	})
}

func main() {}
