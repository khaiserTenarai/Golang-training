package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 9. Go types corresponding to standard GraphQL Scalar types:
// ID -> string, String -> string, Int -> int, Float -> float64, Boolean -> bool
type EmployeeScalars struct {
	ID         string  `json:"id"`         // ID scalar
	Name       string  `json:"name"`       // String scalar
	Age        int     `json:"age"`        // Int scalar
	Salary     float64 `json:"salary"`     // Float scalar
	IsFullTime bool    `json:"isFullTime"` // Boolean scalar
}

type ScalarResponse struct {
	Data struct {
		EmployeeDetails EmployeeScalars `json:"employeeDetails"`
	} `json:"data"`
}

func scalarTypesHandler(w http.ResponseWriter, r *http.Request) {
	emp := EmployeeScalars{
		ID:         "emp-99",
		Name:       "Sarah Connor",
		Age:        34,
		Salary:     87500.50,
		IsFullTime: true,
	}

	resp := ScalarResponse{}
	resp.Data.EmployeeDetails = emp

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/scalars", scalarTypesHandler)
	fmt.Println("Task 9 Server running at http://localhost:8089/graphql/scalars")
	http.ListenAndServe(":8089", nil)
}
