package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee4 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func employeeDetailsHandler(w http.ResponseWriter, r *http.Request) {
	employee := Employee4{
		ID:   101,
		Name: "Abirami",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}
func main() {
	http.HandleFunc("/employee", employeeDetailsHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
