package main

import (
	"encoding/json"
	"log"
	"net/http"
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

var employees = []Employee{
	{ID: 1, Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000},
	{ID: 2, Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000},
	{ID: 3, Name: "Sara Khan", Email: "sara@example.com", Department: "Engineering", Salary: 91000},
	{ID: 4, Name: "Vikram Singh", Email: "vikram@example.com", Department: "Human Resources", Salary: 58000},
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return nil
	}
	return strings.Split(path, "/")
}

func parseID(raw string) (int, bool) {
	id, err := strconv.Atoi(raw)
	if err != nil || id < 1 {
		return 0, false
	}
	return id, true
}

func findByID(id int) (Employee, bool) {
	for _, e := range employees {
		if e.ID == id {
			return e, true
		}
	}
	return Employee{}, false
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	parts := splitPath(r.URL.Path)
	if len(parts) != 2 {
		writeError(w, http.StatusNotFound, "route not found")
		return
	}

	id, ok := parseID(parts[1])
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid employee id")
		return
	}

	e, found := findByID(id)
	if !found {
		writeError(w, http.StatusNotFound, "employee not found")
		return
	}
	writeJSON(w, http.StatusOK, e)
}

func departmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	parts := splitPath(r.URL.Path)

	if len(parts) < 3 || parts[2] != "employee" || len(parts) > 4 {
		writeError(w, http.StatusNotFound, "route not found")
		return
	}

	department := parts[1]
	var inDepartment []Employee
	for _, e := range employees {
		if strings.EqualFold(e.Department, department) {
			inDepartment = append(inDepartment, e)
		}
	}

	if len(parts) == 3 {
		if inDepartment == nil {
			inDepartment = []Employee{}
		}
		writeJSON(w, http.StatusOK, inDepartment)
		return
	}

	id, ok := parseID(parts[3])
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid employee id")
		return
	}
	for _, e := range inDepartment {
		if e.ID == id {
			writeJSON(w, http.StatusOK, e)
			return
		}
	}
	writeError(w, http.StatusNotFound, "employee not found in this department")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee/", employeeHandler)
	mux.HandleFunc("/department/", departmentHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
