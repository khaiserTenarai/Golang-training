/*
Problem Statement:
---------------------
Create a basic HTTP server using Go.

Requirement:
-----------
- Start an HTTP server.
- Server should run on port 8080.
- When the user opens http://localhost:8080,
  the server should display a simple message.

Concept:
-----------
net/http is Go's standard package for creating HTTP servers.
ListenAndServe() starts the server.
*/

package main

import (
	"fmt"
	"net/http"
)

// handler function
func homeHandler(w http.ResponseWriter, r *http.Request) {

	/*
		w -> used to send response to the client.
		r -> contains information about the HTTP request.
	*/

	fmt.Fprintln(w, "Welcome to Go HTTP Server")
}

func main() {

	/*
		Connect "/" URL with homeHandler function.

		Example:
		http://localhost:8080/
	*/
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server started on port 8080")

	/*
		Start the HTTP server.

		:8080 -> server runs on port 8080
		nil   -> use Go's default HTTP router
	*/
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Output in terminal:
--------------------
Server started on port 8080

Open browser:
----------------------------
http://localhost:8080/

Browser output:
---------------------
Welcome to Go HTTP Server

*/