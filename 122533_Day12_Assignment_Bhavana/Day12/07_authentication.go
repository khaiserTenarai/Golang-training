// Day 12 - Task 7: Implement authentication
//
// Authentication answers "who are you?". Here callers have to send
// the secret token you type in when the program starts:
//   curl http://localhost:8080/employee -H "Authorization: Bearer <your token>"
// No token or the wrong one gets a 401. /public needs nothing.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func authMiddleware(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+token {
			w.Header().Set("WWW-Authenticate", "Bearer")
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "missing or wrong token"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	token := ""
	for token == "" {
		token = ask(reader, "Choose a secret token for the API: ")
	}
	employees := sampleEmployees(askCount(reader))
	port := askPort(reader)

	mux := http.NewServeMux()
	mux.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"message": "anyone can see this"})
	})
	mux.Handle("/employee", authMiddleware(token, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, employees)
	})))

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
