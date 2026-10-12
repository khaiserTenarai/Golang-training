package main

import (
	"net/http"
	"strconv"
	"strings"
	"fmt"
)

func main() {

	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idString := strings.TrimPrefix(r.URL.Path, "/employee/")

		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid employee ID", http.StatusBadRequest)
			return
		}

		if id != 1 {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"message":"Employee deleted successfully"}`))
	})
	fmt.Println("http://localhost:8080/employee")
	http.ListenAndServe(":8080", nil)
}