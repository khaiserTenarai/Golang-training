package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type PaginatedResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int   `json:"total_items"`
	Data       []int `json:"data"`
}

func paginateHandler(w http.ResponseWriter, r *http.Request) {
	dataset := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 3 // Default limit
	}

	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex >= len(dataset) {
		json.NewEncoder(w).Encode(PaginatedResponse{Page: page, Limit: limit, TotalItems: len(dataset), Data: []int{}})
		return
	}

	if endIndex > len(dataset) {
		endIndex = len(dataset)
	}

	slicedData := dataset[startIndex:endIndex]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PaginatedResponse{
		Page:       page,
		Limit:      limit,
		TotalItems: len(dataset),
		Data:       slicedData,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /items", paginateHandler)
	http.ListenAndServe(":8080", mux)
}