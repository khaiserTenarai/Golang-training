package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var employees = []Employee{
	{1, "Muneera", "IT", 50000},
	{2, "Rahul", "HR", 45000},
	{3, "Priya", "Finance", 60000},
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {

	// Get ID from URL
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")
	id, _ := strconv.Atoi(idString)

	// Find employee
	for i, employee := range employees {

		if employee.ID == id {

			// Remove employee from slice
			employees = append(employees[:i], employees[i+1:]...)

			fmt.Fprintln(w, "Employee deleted")
			fmt.Println(employees)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/employee/", deleteEmployee)

	fmt.Println("Server running on http://localhost:4040")

	http.ListenAndServe(":4040", nil)
}
