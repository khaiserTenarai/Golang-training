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

var (
	mu        sync.Mutex
	employees = []Employee{
		{ID: 1, Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000},
		{ID: 2, Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000},
	}
	nextID = 3
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
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
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}

		mu.Lock()
		defer mu.Unlock()
		e.ID = nextID
		nextID++
		employees = append(employees, e)
		writeJSON(w, http.StatusCreated, e)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func employeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employee/"))
	if err != nil {
		http.Error(w, "invalid employee id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	index := -1
	for i, e := range employees {
		if e.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		http.Error(w, "employee not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, employees[index])
	case http.MethodPut:
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "invalid JSON body", http.StatusBadRequest)
			return
		}
		e.ID = id
		employees[index] = e
		writeJSON(w, http.StatusOK, e)
	case http.MethodDelete:
		employees = append(employees[:index], employees[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employeesHandler)
	mux.HandleFunc("/employee/", employeeByIDHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
