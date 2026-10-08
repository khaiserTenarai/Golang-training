package middleware

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs every request.
func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log method and URL.
		log.Printf("%s %s", r.Method, r.URL.Path)

		// Continue to the next handler.
		next.ServeHTTP(w, r)
	})
}
