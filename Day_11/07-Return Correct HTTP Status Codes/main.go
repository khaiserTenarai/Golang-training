package main

import (
	"fmt"
	"net/http"
)

func employeeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		// 405 = Method Not Allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 200 = OK
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Employees fetched successfully")
}

func main() {

	http.HandleFunc("/employee", employeeHandler)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		// 500 is normally returned by the server/framework
		// when an internal server error occurs.
		fmt.Println("Server error:", err)
	}
}

/*
Test:

GET:
http://localhost:8080/employee

Response:
Status: 200 OK

Output:
Employees fetched successfully


Try POST:

POST:
http://localhost:8080/employee

Response:
Status: 405 Method Not Allowed


Important:

w.WriteHeader(http.StatusOK)
-> Sends 200 status.

http.Error()
-> Sends an error response with the specified status code.

Common REST API codes:

GET successful       -> 200
POST successful      -> 201
DELETE successful    -> 204
Invalid request      -> 400
Not found            -> 404
Wrong HTTP method    -> 405
*/