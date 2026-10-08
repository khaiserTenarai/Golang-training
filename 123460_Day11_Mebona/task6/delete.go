package main

import (
	"net/http"
	"strconv"
	"strings"
)

var mockDB = map[int]string{
	1: "Alice Smith",
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	if _, exists := mockDB[id]; !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	delete(mockDB, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/employee/", deleteEmployee)
	http.ListenAndServe(":8080", nil)
}