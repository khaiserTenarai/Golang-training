/*
Problem Statement:
Create a GET /employee endpoint.

Requirement:
- Start an HTTP server on port 8080.
- Create /employee endpoint.
- When GET /employee is called,
  return all employees as JSON.

Example:

GET http://localhost:8080/employee

Expected response:

[
    {
        "id": 1,
        "name": "Ganesh",
        "age": 22,
        "email": "ganesh@gmail.com"
    },
    {
        "id": 2,
        "name": "Ravi",
        "age": 23,
        "email": "ravi@gmail.com"
    }
]

Important concepts:
- net/http
- HTTP GET
- http.HandleFunc()
- JSON
- json.NewEncoder()
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

// GET /employee
func getEmployees(w http.ResponseWriter, r *http.Request) {

	/*
		Check whether the request method is GET.

		We only want this endpoint to handle GET requests.
	*/
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	/*
		Set response type as JSON.

		The client will understand that
		the response contains JSON data.
	*/
	w.Header().Set("Content-Type", "application/json")

	/*
		Convert employees slice into JSON
		and send it to the client.

		json.NewEncoder(w).Encode()
		automatically converts Go data into JSON.
	*/
	json.NewEncoder(w).Encode(employees)
}

func main() {

	/*
		Connect /employee URL with getEmployees function.
	*/
	http.HandleFunc("/employee", getEmployees)

	fmt.Println("Server started on port 8080")

	/*
		Start HTTP server.
	*/
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
How to run:

go run main.go

Then open:

http://localhost:8080/employee

Expected output:

[
    {
        "id": 1,
        "name": "Ganesh",
        "age": 22,
        "email": "ganesh@gmail.com"
    },
    {
        "id": 2,
        "name": "Ravi",
        "age": 23,
        "email": "ravi@gmail.com"
    }
]

Interview Questions:

1. What is net/http?
   -> It is Go's standard package for building HTTP clients and servers.

2. What does HandleFunc() do?
   -> It connects a URL path with a handler function.

3. Why do we use json.NewEncoder()?
   -> It converts Go data into JSON and writes it directly to the response.

4. Why do we set Content-Type?
   -> It tells the client that the response is JSON.
*/