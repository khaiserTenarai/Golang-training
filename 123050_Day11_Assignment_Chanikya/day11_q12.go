package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		log.Printf(
			"METHOD=%s PATH=%s TIME=%v",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Employee API")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/employee", employeeHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(
		":8080",
		loggingMiddleware(mux),
	)
}
