package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func askPort(reader *bufio.Reader) string {
	port := ask(reader, "Port to run the server on (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}
	return port
}

func askCount(reader *bufio.Reader) int {
	count, err := strconv.Atoi(ask(reader, "How many sample employees should be generated? "))
	if err != nil || count < 0 {
		fmt.Println("That's not a valid number, generating 10.")
		count = 10
	}
	return count
}

// sampleEmployees makes up `count` employees spread over a few departments.
func sampleEmployees(count int) []Employee {
	departments := []string{"IT", "HR", "Finance", "Sales"}
	list := make([]Employee, 0, count)
	for i := 1; i <= count; i++ {
		list = append(list, Employee{
			ID:         i,
			Name:       fmt.Sprintf("Employee %d", i),
			Email:      fmt.Sprintf("employee%d@example.com", i),
			Department: departments[i%len(departments)],
			Salary:     float64(30000 + (i*7919)%40000),
		})
	}
	return list
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

type contextKey string

const roleKey contextKey = "role"

var (
	employees []Employee
	mu        sync.Mutex
)

func authenticate(tokens map[string]string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		role, ok := tokens[token]
		if !ok {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing or wrong token"})
			return
		}
		ctx := context.WithValue(r.Context(), roleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func requireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, _ := r.Context().Value(roleKey).(string)
		if role != "admin" {
			writeJSON(w, http.StatusForbidden, map[string]string{"error": "only admins can do this"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func listEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	writeJSON(w, http.StatusOK, employees)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "use DELETE here"})
		return
	}
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a number"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "employee not found"})
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	adminToken := ""
	for adminToken == "" {
		adminToken = ask(reader, "Choose an ADMIN token: ")
	}
	userToken := ""
	for userToken == "" || userToken == adminToken {
		userToken = ask(reader, "Choose a different USER token: ")
	}
	employees = sampleEmployees(askCount(reader))
	port := askPort(reader)

	tokens := map[string]string{adminToken: "admin", userToken: "user"}

	mux := http.NewServeMux()
	mux.Handle("/employee", authenticate(tokens, http.HandlerFunc(listEmployees)))
	mux.Handle("/employee/", authenticate(tokens, requireAdmin(http.HandlerFunc(deleteEmployee))))

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
