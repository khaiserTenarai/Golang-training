package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var (
	mu        sync.Mutex
	employees = map[int]Employee{
		1: {ID: 1, Name: "Indu", Position: "Developer"},
	}
)

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	mu.Lock()
	_, exists := employees[id]
	if !exists {
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})
		return
	}

	delete(employees, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /employee/{id}", deleteEmployeeHandler)
	http.ListenAndServe(":8080", mux)
}