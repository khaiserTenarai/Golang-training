// Bonus: Recovery Middleware — catches panics and returns 500 instead of crashing.

package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[PANIC] %v — %s %s", err, r.Method, r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"message": "Internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
