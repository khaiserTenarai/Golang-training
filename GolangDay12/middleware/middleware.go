package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

func RequestIDAndLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := "req-" + time.Now().Format("20060102150405") // 5. Request-ID[cite: 2]
		ctx := context.WithValue(r.Context(), "RequestID", reqID)
		
		log.Printf("[INFO] %s - %s %s", reqID, r.Method, r.URL.Path) // 6. Logging[cite: 2]
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized - Missing Token", http.StatusUnauthorized) // 7. Authentication[cite: 2]
			return
		}
		
		// 8. Authorization check (dummy example: only "admin-token" can POST)[cite: 2]
		if r.Method == http.MethodPost && token != "Bearer admin-token" {
			http.Error(w, "Forbidden - Insufficient Permissions", http.StatusForbidden)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}