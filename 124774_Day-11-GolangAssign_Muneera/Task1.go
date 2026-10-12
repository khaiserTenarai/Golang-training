package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Employee API is running!")
}

func main() {

	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on http://localhost:1010")

	err := http.ListenAndServe(":1010", nil)

	if err != nil {
		log.Fatal(err)
	}
}
