package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Employee struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Email        string      `json:"email"`
	Role         string      `json:"role"`
	Salary       float64     `json:"salary"`
	IsActive     bool        `json:"isActive"`
	DepartmentID string      `json:"-"`
	Department   *Department `json:"department,omitempty"`
}

type CreateEmployeeInput struct {
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Role         string  `json:"role"`
	Salary       float64 `json:"salary"`
	DepartmentID string  `json:"departmentId"`
}

type GraphQLError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type Response struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

var (
	mutex       sync.Mutex
	departments = []Department{
		{ID: "d-1", Name: "Technology", Location: "Floor 3"},
		{ID: "d-2", Name: "Operations", Location: "Floor 1"},
	}
	employees = []Employee{
		{ID: "e-1", Name: "Alice Smith", Email: "alice@co.com", Role: "Tech Lead", Salary: 110000.0, IsActive: true, DepartmentID: "d-1"},
	}
)

func findDepartment(id string) *Department {
	for _, d := range departments {
		if d.ID == id {
			return &d
		}
	}
	return nil
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()

	switch r.Method {
	case http.MethodGet:
		// Resolve list query & field resolvers
		var resolvedEmployees []Employee
		for _, e := range employees {
			e.Department = findDepartment(e.DepartmentID)
			resolvedEmployees = append(resolvedEmployees, e)
		}
		json.NewEncoder(w).Encode(Response{Data: map[string]interface{}{"employees": resolvedEmployees}})

	case http.MethodPost:
		var input CreateEmployeeInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			json.NewEncoder(w).Encode(Response{Errors: []GraphQLError{{Message: "Invalid input", Code: "BAD_USER_INPUT"}}})
			return
		}

		if findDepartment(input.DepartmentID) == nil {
			json.NewEncoder(w).Encode(Response{Errors: []GraphQLError{{Message: "Department not found", Code: "BAD_USER_INPUT"}}})
			return
		}

		newEmp := Employee{
			ID:           fmt.Sprintf("e-%d", len(employees)+1),
			Name:         input.Name,
			Email:        input.Email,
			Role:         input.Role,
			Salary:       input.Salary,
			IsActive:     true,
			DepartmentID: input.DepartmentID,
		}
		newEmp.Department = findDepartment(newEmp.DepartmentID)
		employees = append(employees, newEmp)

		json.NewEncoder(w).Encode(Response{Data: map[string]interface{}{"createEmployee": newEmp}})

	case http.MethodDelete:
		id := r.URL.Query().Get("id")
		deleted := false
		for i, e := range employees {
			if e.ID == id {
				employees = append(employees[:i], employees[i+1:]...)
				deleted = true
				break
			}
		}
		if !deleted {
			json.NewEncoder(w).Encode(Response{Errors: []GraphQLError{{Message: "Employee not found", Code: "NOT_FOUND"}}})
			return
		}
		json.NewEncoder(w).Encode(Response{Data: map[string]bool{"deleteEmployee": true}})
	}
}

func main() {
	http.HandleFunc("/api/graphql", apiHandler)
	fmt.Println("Task 15 Mini Project Server running at http://localhost:8095/api/graphql")
	http.ListenAndServe(":8095", nil)
}
