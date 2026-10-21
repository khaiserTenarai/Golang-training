package middleware

import (
	"log"
	"net/http"
	"time"
)

// statusWriter captures the HTTP response status for logging.
type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// Logging records request method, path, status, and duration.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sw, r)

		log.Printf("method=%s path=%s status=%d duration=%s",
			r.Method, r.URL.Path, sw.status, time.Since(start))
	})
}
