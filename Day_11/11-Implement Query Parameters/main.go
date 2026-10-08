package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var employees = []Employee{
	{ID: 1, Name: "Ganesh", Age: 22, Email: "ganesh@gmail.com"},
	{ID: 2, Name: "Ravi", Age: 23, Email: "ravi@gmail.com"},
	{ID: 3, Name: "Ganesh Kumar", Age: 25, Email: "ganeshkumar@gmail.com"},
}

func getEmployees(w http.ResponseWriter, r *http.Request) {

	// Get name from query parameter.
	name := r.URL.Query().Get("name")

	var result []Employee

	for _, employee := range employees {

		if strings.EqualFold(employee.Name, name) {
			result = append(result, employee)
		}
	}

	if len(result) == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}

func main() {

	http.HandleFunc("/employee", getEmployees)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Test:

GET:
http://localhost:8080/employee?name=Ganesh

Response:

[
    {
        "id": 1,
        "name": "Ganesh",
        "age": 22,
        "email": "ganesh@gmail.com"
    }
]

Important:

URL:
 /employee?name=Ganesh
            ↑
      Query parameter

r.URL.Query().Get("name")
        ↓
      "Ganesh"

Path parameter:
    /employee/1

Query parameter:
    /employee?name=Ganesh
*/