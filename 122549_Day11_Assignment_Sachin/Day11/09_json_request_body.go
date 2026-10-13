package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		mu.Lock()
		defer mu.Unlock()
		writeJSON(w, http.StatusOK, employees)
		return
	}

	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		writeJSON(w, http.StatusUnsupportedMediaType, map[string]string{"error": "Content-Type must be application/json"})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB is plenty for one employee
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var e Employee
	if err := decoder.Decode(&e); err != nil {
		message := "the body is not valid JSON: " + err.Error()
		if errors.Is(err, io.EOF) {
			message = "the request body is empty"
		}
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": message})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	e.ID = nextID
	nextID++
	employees = append(employees, e)
	writeJSON(w, http.StatusCreated, e)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	port := askPort(reader)

	http.HandleFunc("/employee", employeesHandler)

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
