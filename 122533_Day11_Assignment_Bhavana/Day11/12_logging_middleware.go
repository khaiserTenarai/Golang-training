package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func askEmployees(reader *bufio.Reader) []Employee {
	count, err := strconv.Atoi(ask(reader, "How many employees do you want to start with? "))
	if err != nil || count < 0 {
		fmt.Println("That's not a valid number, starting with no employees.")
		count = 0
	}

	list := []Employee{}
	for i := 1; i <= count; i++ {
		fmt.Printf("\n--- Employee %d ---\n", i)
		name := ask(reader, "Name: ")
		email := ask(reader, "Email: ")
		department := ask(reader, "Department: ")
		salary, _ := strconv.ParseFloat(ask(reader, "Salary: "), 64)
		list = append(list, Employee{ID: i, Name: name, Email: email, Department: department, Salary: salary})
	}
	return list
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// statusRecorder remembers the status code, because a plain
// ResponseWriter doesn't let us read it back afterwards.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(recorder, r)

		log.Printf("%s %s -> %d (%v)", r.Method, r.URL.Path, recorder.status, time.Since(start))
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	employees := askEmployees(reader)
	port := askPort(reader)

	mux := http.NewServeMux()

	mux.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, employees)
	})

	mux.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be a number"})
			return
		}
		for _, e := range employees {
			if e.ID == id {
				writeJSON(w, http.StatusOK, e)
				return
			}
		}
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "employee not found"})
	})

	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(":"+port, loggingMiddleware(mux)); err != nil {
		fmt.Println("Server stopped:", err)
	}
}
