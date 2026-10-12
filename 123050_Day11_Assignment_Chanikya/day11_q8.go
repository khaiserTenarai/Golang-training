package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func responseHandler(w http.ResponseWriter, r *http.Request) {

	response := Response{
		Message: "Employee retrieved successfully",
		Status:  "success",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/employee", responseHandler)

	http.ListenAndServe(":8080", nil)
}
