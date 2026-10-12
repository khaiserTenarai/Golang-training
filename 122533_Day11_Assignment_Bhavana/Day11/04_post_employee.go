package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

var (
	employees = []Employee{}
	nextID    = 1
	mu        sync.Mutex
)

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(employees)

	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "could not read the JSON body", http.StatusBadRequest)
			return
		}

		mu.Lock()
		e.ID = nextID
		nextID++
		employees = append(employees, e)
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	port := askPort(reader)

	http.HandleFunc("/employee", employeesHandler)

	fmt.Println("Server running on port", port, "- send a POST to /employee")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
