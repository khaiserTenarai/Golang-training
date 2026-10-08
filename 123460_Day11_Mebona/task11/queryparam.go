package main

import (
	"fmt"
	"net/http"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	queryVals := r.URL.Query()

	name := queryVals.Get("name")
	position := queryVals.Get("position")
	limit := queryVals.Get("limit")

	fmt.Fprintf(w, "Query Parameters Received:\nName: %s\nPosition: %s\nLimit: %s\n", name, position, limit)
}

func main() {
	http.HandleFunc("/search", queryParamHandler)
	http.ListenAndServe(":8080", nil)
}