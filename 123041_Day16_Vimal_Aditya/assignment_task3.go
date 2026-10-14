package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type GraphQLResponse struct {
	Data struct {
		SampleDepartment Department `json:"sampleDepartment"`
	} `json:"data"`
}

func departmentHandler(w http.ResponseWriter, r *http.Request) {
	dept := Department{ID: "dept-1", Name: "Engineering", Location: "Building A"}
	res := GraphQLResponse{}
	res.Data.SampleDepartment = dept

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/graphql", departmentHandler)
	fmt.Println("Server running at http://localhost:8083/graphql")
	http.ListenAndServe(":8083", nil)
}
