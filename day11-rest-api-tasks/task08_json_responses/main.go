package main

import (
	"encoding/json"
	"log"
	"net/http"
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

type Response struct {
	Success bool           `json:"success"`
	Data    any            `json:"data,omitempty"`
	Error   string         `json:"error,omitempty"`
	Meta    map[string]any `json:"meta,omitempty"`
}

var (
	mu        sync.RWMutex
	employees = []Employee{
		{ID: 1, Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000},
		{ID: 2, Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000},
		{ID: 3, Name: "Sara Khan", Email: "sara@example.com", Department: "Engineering", Salary: 91000},
	}
	nextID = 4
)

func respond(w http.ResponseWriter, status int, resp Response) {
	body, err := json.Marshal(resp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success":false,"error":"failed to encode response"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(body)
}

func respondOK(w http.ResponseWriter, status int, data any) {
	respond(w, status, Response{Success: true, Data: data})
}

func respondError(w http.ResponseWriter, status int, message string) {
	respond(w, status, Response{Success: false, Error: message})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotFound, "route not found: "+r.URL.Path)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	respondOK(w, http.StatusOK, map[string]string{"status": "ok"})
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
		respond(w, http.StatusOK, Response{
			Success: true,
			Data:    employees,
			Meta:    map[string]any{"count": len(employees)},
		})
	case http.MethodPost:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			respondError(w, http.StatusBadRequest, "invalid JSON body")
			return
		}
		mu.Lock()
		defer mu.Unlock()
		e.ID = nextID
		nextID++
		employees = append(employees, e)
		respondOK(w, http.StatusCreated, e)
	default:
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func employeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	mu.RLock()
	defer mu.RUnlock()
	for _, e := range employees {
		if e.ID == id {
			respondOK(w, http.StatusOK, e)
			return
		}
	}
	respondError(w, http.StatusNotFound, "employee not found")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFoundHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/employee", employeesHandler)
	mux.HandleFunc("/employee/", employeeByIDHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
