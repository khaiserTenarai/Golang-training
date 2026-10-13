package main

import (
	"fmt"
	"net/http"
)

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE method is allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Employee deleted successfully")
}
func main() {
	http.HandleFunc("/employee/delete", deleteEmployeeHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
