// Task 6: Logging Middleware
// Logs method, path, status code, duration, and request ID for every request.

package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := &statusWriter{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(lw, r)

		reqID := GetRequestID(r)
		log.Printf("[%s] %s %s → %d (%v) [%s]",
			r.Method, r.URL.Path, r.RemoteAddr,
			lw.statusCode, time.Since(start).Round(time.Microsecond), reqID)
	})
}

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
