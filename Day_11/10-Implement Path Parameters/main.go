package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
}

func getEmployee(w http.ResponseWriter, r *http.Request) {

	// Example: /employee/1
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")

	// Convert path parameter from string to int.
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	for _, employee := range employees {

		if employee.ID == id {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)

			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/employee/", getEmployee)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Test:

GET http://localhost:8080/employee/1

Response:

{
    "id": 1,
    "name": "Ganesh",
    "age": 22,
    "email": "ganesh@gmail.com"
}


GET http://localhost:8080/employee/2

Response:

{
    "id": 2,
    "name": "Ravi",
    "age": 23,
    "email": "ravi@gmail.com"
}


Important:

/employee/1
          ↑
     Path parameter

r.URL.Path
    ↓
/employee/1

strings.TrimPrefix()
    ↓
1

strconv.Atoi()
    ↓
1 (integer)

Then search employee with ID 1.
*/