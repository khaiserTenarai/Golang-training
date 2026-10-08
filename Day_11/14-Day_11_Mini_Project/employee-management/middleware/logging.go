// Package middleware contains HTTP middleware components.
package middleware

// Import packages required for HTTP logging.
import (
	// Import fmt for formatting log messages.
	"fmt"
	// Import net/http for HTTP server types.
	"net/http"
	// Import time for measuring request duration.
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
		// Create a response writer wrapper.
		wrappedWriter := &ResponseWriter{ResponseWriter: w, status: http.StatusOK}
		// Execute the next handler.
		next.ServeHTTP(wrappedWriter, r)
		// Calculate request duration.
		duration := time.Since(start)
		// Write the request information to the application logger.
		logger.Info(fmt.Sprintf("%s %s %d %s", r.Method, r.URL.RequestURI(), wrappedWriter.status, duration))
	})
}
