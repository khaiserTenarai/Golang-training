package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		//parse query parameters from the url
		name := r.URL.Query().Get("name")
		page := r.URL.Query().Get("page")
		// Return the values to deomostarte parsing
		fmt.Fprintln(w, "name:", name, "page:", page)
	})
	// Start the server.
	fmt.Println("Server runnig at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

//Request url -http://localhost:8080/employees?name=muneera&page=10
