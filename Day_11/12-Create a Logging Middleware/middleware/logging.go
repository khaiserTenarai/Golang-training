package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs each request.
func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log HTTP method and URL.
		log.Printf("%s %s", r.Method, r.URL.Path)

		// Call the next handler.
		next.ServeHTTP(w, r)
	})
}
