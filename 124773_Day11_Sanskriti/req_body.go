package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		var e Employee

		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "invalid JSON", 400)
			return
		}

		fmt.Fprintln(w, "Received:", e.Name, e.Email)
	})
	fmt.Println("connection success http://localhost:8080/employees")
	http.ListenAndServe(":8080", nil)
	
}