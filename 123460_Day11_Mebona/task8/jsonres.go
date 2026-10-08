package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  "success",
		Message: "Data retrieved successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/json", jsonResponseHandler)
	http.ListenAndServe(":8080", nil)
}