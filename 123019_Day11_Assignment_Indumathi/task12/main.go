package main

import (
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		
		// Pass request down the chain
		next.ServeHTTP(w, r)

		log.Printf("[%s] %s %s - %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(startTime))
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from handler!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", mainHandler)

	wrappedMux := loggingMiddleware(mux)

	log.Println("Starting server with logging middleware on :8080...")
	http.ListenAndServe(":8080", wrappedMux)
}