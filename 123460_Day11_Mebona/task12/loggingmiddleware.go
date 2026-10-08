package main

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware logs the incoming HTTP method, path, and duration taken
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		fmt.Printf("[%s] %s %s - Duration: %v\n", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Middleware!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	// Wrap mux with logging middleware
	loggedMux := LoggingMiddleware(mux)

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", loggedMux)
}