package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GraphQLError struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

type ErrorResponse struct {
	Data   interface{}    `json:"data"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json")

	if id != "1" {
		resp := ErrorResponse{
			Data: nil,
			Errors: []GraphQLError{
				{
					Message: fmt.Sprintf("Employee with ID '%s' was not found.", id),
					Extensions: map[string]interface{}{
						"code": "EMPLOYEE_NOT_FOUND",
						"http": map[string]int{"status": 404},
					},
				},
			},
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": map[string]interface{}{
			"employee": map[string]string{"id": "1", "name": "John Doe"},
		},
	})
}

func main() {
	http.HandleFunc("/graphql/error", errorHandler)
	fmt.Println("Server running at http://localhost:8080/graphql/error?id=99")
	http.ListenAndServe(":8080", nil)
}
