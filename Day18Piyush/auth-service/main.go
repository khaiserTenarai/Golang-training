// Task 6: Define Authentication Service
// Task 7: Database-per-service (own user store)
// Run: go run main.go
// Port: 8083

package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Token    string `json:"token,omitempty"`
}

var (
	mu     sync.RWMutex
	users  = map[string]User{}
	tokens = map[string]User{}
	next   = 1
)

func init() {
	seed := []User{
		{Username: "admin", Password: "admin123", Role: "admin"},
		{Username: "piyush", Password: "piyush123", Role: "user"},
		{Username: "manager", Password: "manager123", Role: "manager"},
	}
	for _, u := range seed {
		u.ID = next
		next++
		users[u.Username] = u
	}
}

func genToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	mu.RLock()
	user, ok := users[req.Username]
	mu.RUnlock()
	if !ok || user.Password != req.Password {
		respond(w, 401, map[string]interface{}{"error": "invalid credentials"})
		return
	}
	token := genToken()
	mu.Lock()
	tokens[token] = user
	mu.Unlock()
	respond(w, 200, map[string]interface{}{
		"token": token, "user_id": user.ID, "username": user.Username, "role": user.Role,
		"expires_at": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	})
}

// Validate token endpoint — called by API Gateway (Task 10)
func validateHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	mu.RLock()
	user, ok := tokens[token]
	mu.RUnlock()
	if !ok {
		respond(w, 401, map[string]interface{}{"valid": false, "error": "invalid token"})
		return
	}
	respond(w, 200, map[string]interface{}{
		"valid": true, "user_id": user.ID, "username": user.Username, "role": user.Role,
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, 200, map[string]interface{}{"service": "auth-service", "status": "healthy", "port": 8083})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/auth/login", loginHandler)
	mux.HandleFunc("/auth/validate", validateHandler)

	fmt.Println("[Auth-Service] Running on http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", mux))
}
