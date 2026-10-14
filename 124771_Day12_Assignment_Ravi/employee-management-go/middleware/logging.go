package middleware

import (
	"fmt"

	"net/http"

	"time"
)

// ResponseWriter wraps http.ResponseWriter so the status code can be captured.
type ResponseWriter struct {
	// ResponseWriter embeds the standard HTTP response writer.
	http.ResponseWriter
	// status stores the HTTP response status code.
	status int
}

// WriteHeader captures the response status before writing it to the client.
func (rw *ResponseWriter) WriteHeader(status int) {
	// Store the status code.
	rw.status = status
	// Call the original WriteHeader implementation.
	rw.ResponseWriter.WriteHeader(status)
}

// Logging returns HTTP request logging middleware.
func Logging(next http.Handler, logger interface{ Info(string) }) http.Handler {
	// Return a handler that wraps the next handler.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the request start time.
		start := time.Now()

		wrappedWriter := &ResponseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(wrappedWriter, r)

		duration := time.Since(start)

		logger.Info(fmt.Sprintf("%s %s %d %s", r.Method, r.URL.RequestURI(), wrappedWriter.status, duration))
	})
}
