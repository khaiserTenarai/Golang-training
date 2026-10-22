package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (rw *ResponseWriter) WriteHeader(status int) {
	rw.Status = status
	rw.ResponseWriter.WriteHeader(status)
}

func Logging(next http.Handler, logger interface{ Info(string) }) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &ResponseWriter{ResponseWriter: w, Status: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		logger.Info(fmt.Sprintf("%s %s %d %s", r.Method, r.URL.RequestURI(), wrapped.Status, time.Since(start)))
	})
}