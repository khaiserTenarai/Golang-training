package main

import (
	"fmt"
	"net/http"
)

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Employee details")
}
func main() {
	http.HandleFunc("/employee", employeeHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
