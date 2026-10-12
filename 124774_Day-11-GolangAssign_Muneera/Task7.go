package main

import (
	"fmt"
	"net/http"
)

func employee(w http.ResponseWriter, r *http.Request) {

	// Successful request
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Employee found")
		return
	}

	// New employee created
	if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Employee created")
		return
	}

	// Invalid request
	if r.Method == "PUT" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Employee not found
	if r.Method == "DELETE" {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
}

func main() {

	http.HandleFunc("/employee", employee)

	fmt.Println("Server running on http://localhost:7171")

	http.ListenAndServe(":7171", nil)
}

//http://localhost:7171/employee-for get,put,post,delete..
