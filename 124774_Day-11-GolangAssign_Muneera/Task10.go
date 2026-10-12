package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {

		// Extract the ID from the URL path
		id := strings.TrimPrefix(r.URL.Path, "/employees/")

		// Return the extracted ID
		fmt.Fprintln(w, "Employee ID:", id)
	})

	// Start the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
