package main

import (
	"fmt"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Employee details")
}
func employeeDetailsHandler(w http.ResponseWriter, r *http.Request) {
	employee := Employee{
		ID:   101,
		Name: "Tom",
	}
	fmt.Fprintln(w, employee)
}
func main() {
	http.HandleFunc("/employee", employeeHandler)
	http.HandleFunc("/employee/detail", employeeDetailsHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
