package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Logging records method, path, and elapsed time.
// Task 15: logging middleware for the Employee API.
func Logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start).String(),
		)
	})
}
