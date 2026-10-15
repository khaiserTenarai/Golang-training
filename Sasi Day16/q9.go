package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EmployeeScalars struct {
	ID         string  `json:"id"`      
	Name       string  `json:"name"`    
	Age        int     `json:"age"`  
	Salary     float64 `json:"salary"` 
	IsFullTime bool    `json:"isFullTime"`
}

type ScalarResponse struct {
	Data struct {
		EmployeeDetails EmployeeScalars `json:"employeeDetails"`
	} `json:"data"`
}

func scalarTypesHandler(w http.ResponseWriter, r *http.Request) {
	emp := EmployeeScalars{
		ID:         "emp1",
		Name:       "Sasi",
		Age:        23,
		Salary:     90000.00,
		IsFullTime: true,
	}

	resp := ScalarResponse{}
	resp.Data.EmployeeDetails = emp

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/scalars", scalarTypesHandler)
	fmt.Println("Server running at http://localhost:8089/graphql/scalars")
	http.ListenAndServe(":8089", nil)
}
