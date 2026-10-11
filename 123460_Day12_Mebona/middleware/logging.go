package middleware

import (
	"fmt"
	"net/http"
	"time"

	"employee-management/logging"
)

func Logging(next http.Handler, logger *logging.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID, _ := r.Context().Value(RequestIDKey).(string)

		next.ServeHTTP(w, r)

		logger.Info(fmt.Sprintf("[%s] %s %s %s", reqID, r.Method, r.URL.Path, time.Since(start)))
	})
}