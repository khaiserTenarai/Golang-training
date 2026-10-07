package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", helloHandler)

	loggedMux := loggingMiddleware(mux)

	http.ListenAndServe(":8080", loggedMux)
}