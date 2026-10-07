package main

import (
	"net/http"
	"strconv"
	"sync"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var (
	employees = map[int]Employee{
		1: {ID: 1, Name: "Lakshmi Shibu", Department: "Engineering"},
	}
	mu sync.Mutex
)

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := employees[id]; !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	delete(employees, id)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("DELETE /employee/{id}", deleteEmployeeHandler)

	http.ListenAndServe(":8080", mux)
}