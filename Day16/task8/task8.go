package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type DeleteResponse struct {
	Data struct {
		DeleteEmployee bool `json:"deleteEmployee"`
	} `json:"data"`
}

var employeesDB = []Employee{
	{ID: "1", Name: "David"},
	{ID: "2", Name: "Emma"},
}

// 8. Mutation for deleting an employee
func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var req DeleteRequest
	json.NewDecoder(r.Body).Decode(&req)

	success := false
	for i, emp := range employeesDB {
		if emp.ID == req.ID {
			employeesDB = append(employeesDB[:i], employeesDB[i+1:]...)
			success = true
			break
		}
	}

	resp := DeleteResponse{}
	resp.Data.DeleteEmployee = success

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/delete", deleteEmployeeHandler)
	fmt.Println("Task 8 Server running at http://localhost:8088/graphql/delete")
	http.ListenAndServe(":8088", nil)
}
