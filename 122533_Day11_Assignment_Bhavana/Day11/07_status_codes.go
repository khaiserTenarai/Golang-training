package main

import (
	"bufio"
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

var (
	employees = []Employee{}
	nextID    = 1
	mu        sync.Mutex
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func findIndex(id int) int {
	for i, e := range employees {
		if e.ID == id {
			return i
		}
	}
	return -1
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		writeJSON(w, http.StatusOK, employees)

	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}

		mu.Lock()
		defer mu.Unlock()
		for _, existing := range employees {
			if strings.EqualFold(existing.Email, e.Email) {
				writeJSON(w, http.StatusConflict, map[string]string{"error": "that email is already in use"})
				return
			}
		}
		e.ID = nextID
		nextID++
		employees = append(employees, e)
		writeJSON(w, http.StatusCreated, e)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a number"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	idx := findIndex(id)
	if idx == -1 {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "employee not found"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, employees[idx])

	case http.MethodPut:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
			return
		}
		e.ID = id
		employees[idx] = e
		writeJSON(w, http.StatusOK, e)

	case http.MethodDelete:
		employees = append(employees[:idx], employees[idx+1:]...)
		w.WriteHeader(http.StatusNoContent)

	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	port := askPort(reader)

	http.HandleFunc("/employee", employeesHandler)
	http.HandleFunc("/employee/", employeeHandler)

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
