package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Handles paths matching /employee/{id}/department/{dept_id}
func pathParamHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) < 4 || parts[0] != "employee" || parts[2] != "department" {
		http.Error(w, "Route not found", http.StatusNotFound)
		return
	}

	empID := parts[1]
	deptID := parts[3]

	fmt.Fprintf(w, "Employee ID: %s, Department ID: %s", empID, deptID)
}

func main() {
	http.HandleFunc("/employee/", pathParamHandler)
	http.ListenAndServe(":8080", nil)
}