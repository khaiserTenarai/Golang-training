package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HTTP method:", r.Method)
	fmt.Println("URL path:", r.URL.Path)
	fmt.Println("Protocol:", r.Proto)
	fmt.Println("User-Agent:", r.Header.Get("User-Agent"))

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Request received successfully")
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server running at http://localhost:10004")

	err := http.ListenAndServe(":10004", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}