package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	DepartmentID string      `json:"-"`
	Department   *Department `json:"department"`
}

var departments = []Department{
	{ID: "d1", Name: "Engineering"},
}

var employees = []Employee{
	{ID: "e1", Name: "Sasi", DepartmentID: "d1"},
}

func resolveDepartment(deptID string) *Department {
	for _, d := range departments {
		if d.ID == deptID {
			return &d
		}
	}
	return nil
}

func resolverHandler(w http.ResponseWriter, r *http.Request) {
	emp := employees[0]
	emp.Department = resolveDepartment(emp.DepartmentID)

	resp := map[string]interface{}{
		"data": map[string]interface{}{
			"employee": emp,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/resolvers", resolverHandler)
	fmt.Println("Server running at http://localhost:8093/graphql/resolvers")
	http.ListenAndServe(":8093", nil)
}
