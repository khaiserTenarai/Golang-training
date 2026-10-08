package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {

	employee := Employee{
		ID:    1,
		Name:  "Ganesh",
		Age:   22,
		Email: "ganesh@gmail.com",
	}

	// Tell the client that the response is JSON.
	w.Header().Set("Content-Type", "application/json")

	// Return 200 OK.
	w.WriteHeader(http.StatusOK)

	// Convert Go struct to JSON.
	json.NewEncoder(w).Encode(employee)
}

func main() {

	http.HandleFunc("/employee", employeeHandler)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Test:

GET http://localhost:8080/employee

Response:

{
    "id": 1,
    "name": "Ganesh",
    "age": 22,
    "email": "ganesh@gmail.com"
}

Important:

w.Header().Set("Content-Type", "application/json")
-> Tells the client the response is JSON.

json.NewEncoder(w).Encode(employee)
-> Converts the Go struct into JSON
   and sends it to the client.

Go struct → JSON
*/