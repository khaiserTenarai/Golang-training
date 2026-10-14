package controller

import (
	"crypto/subtle"
	"log/slog"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func LogRequests(log *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sw, r)

		log.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", sw.status,
			"took", time.Since(start).String(),
		)
	})
}

func requireKey(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			given := r.Header.Get("X-API-Key")
			if subtle.ConstantTimeCompare([]byte(given), []byte(apiKey)) != 1 {
				writeError(w, http.StatusUnauthorized, "missing or wrong API key")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
