package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const (
	requestIDKey contextKey = "requestID"
	roleKey      contextKey = "role"
)

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func newRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "unknown"
	}
	return hex.EncodeToString(b)
}

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-ID")
		if id == "" {
			id = newRequestID()
		}
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(recorder, r)

		id, _ := r.Context().Value(requestIDKey).(string)
		log.Printf("[%s] %s %s -> %d (%v)", id, r.Method, r.URL.Path, recorder.status, time.Since(start))
	})
}

type Auth struct {
	roles map[string]string // token -> role
}

func NewAuth(adminToken, userToken string) *Auth {
	return &Auth{roles: map[string]string{adminToken: "admin", userToken: "user"}}
}

// Authenticate works out who is calling (401 if we can't).
func (a *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		role, ok := a.roles[token]
		if !ok {
			writeError(w, http.StatusUnauthorized, "missing or invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), roleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Auth) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(roleKey).(string)
		if r.Method != http.MethodGet && role != "admin" {
			writeError(w, http.StatusForbidden, "only admins can change employee data")
			return
		}
		next.ServeHTTP(w, r)
	})
}
