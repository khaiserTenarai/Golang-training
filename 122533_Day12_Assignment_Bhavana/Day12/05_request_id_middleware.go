// Day 12 - Task 5: Create request-ID middleware
//
// Every request gets a unique ID. If the caller already sent an
// X-Request-ID header we keep theirs, otherwise we make one. The ID is
// put on the response header and into the request context, so any
// handler further down can read it (handy when you're hunting through
// logs for one particular request).
//   curl -i http://localhost:8080/employee
package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func newRequestID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "unknown"
	}
	return hex.EncodeToString(b)
}

func requestIDMiddleware(next http.Handler) http.Handler {
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What message should /employee reply with? ")
	messageText, _ := reader.ReadString('\n')
	message := strings.TrimSpace(messageText)

	fmt.Print("Port to run the server on (press Enter for 8080): ")
	portText, _ := reader.ReadString('\n')
	port := strings.TrimSpace(portText)
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		id, _ := r.Context().Value(requestIDKey).(string)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message":    message,
			"request_id": id,
		})
	})

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, requestIDMiddleware(mux)); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
