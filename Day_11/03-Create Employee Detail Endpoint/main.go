/*
Problem Statement:
Create an employee detail endpoint.

Requirement:
- GET /employee/{id}
- Get employee details using ID.
- If employee exists, return employee as JSON.
- If employee does not exist, return 404.

Examples:

GET /employee/1
GET /employee/2

Important concepts:
- URL path
- strconv.Atoi()
- HTTP status codes
- JSON response
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Employee represents employee details.
type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// Employee data stored in memory.
var employees = []Employee{
	{
		ID:    1,
		Name:  "Ganesh",
		Age:   22,
		Email: "ganesh@gmail.com",
	},
	{
		ID:    2,
		Name:  "Ravi",
		Age:   23,
		Email: "ravi@gmail.com",
	},
}

// GET /employee/{id}
func getEmployeeByID(w http.ResponseWriter, r *http.Request) {

	/*
		Example URL:

		/employee/1

		r.URL.Path gives:

		/employee/1
	*/

	path := r.URL.Path

	/*
		Remove "/employee/" from the URL.

		/employee/1
		        ↓
		1
	*/
	idString := strings.TrimPrefix(path, "/employee/")

	/*
		strconv.Atoi() converts string to integer.

		"1" -> 1
	*/
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	/*
		Search employee by ID.
	*/
	for _, employee := range employees {

		if employee.ID == id {

			/*
				Employee found.

				Set response as JSON.
			*/
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(employee)

			return
		}
	}

	/*
		If the loop finishes without finding
		the employee, return 404.
	*/
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	/*
		Connect /employee/ with the handler.

		The trailing slash allows paths such as:

		/employee/1
		/employee/2
	*/
	http.HandleFunc("/employee/", getEmployeeByID)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
How to run:

go run main.go

Test:

GET http://localhost:8080/employee/1

Expected:

{
    "id": 1,
    "name": "Ganesh",
    "age": 22,
    "email": "ganesh@gmail.com"
}

Test another employee:

GET http://localhost:8080/employee/2

If employee doesn't exist:

GET http://localhost:8080/employee/10

Response:

Employee not found

HTTP status:

404 Not Found


Important Interview Points:

1. What is r.URL.Path?
   -> It gives the path portion of the requested URL.

2. Why use strconv.Atoi()?
   -> URL values are strings, so we convert the employee ID
      from string to integer.

3. What is HTTP 404?
   -> It means the requested resource was not found.

4. What is HTTP 400?
   -> It means the client sent an invalid request.

5. Why do we use return after finding the employee?
   -> To stop the function immediately after sending the response.
*/