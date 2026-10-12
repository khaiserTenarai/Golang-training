package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{Name: "Laptop", Price: 1200},
	{Name: "Mouse", Price: 25},
	{Name: "Keyboard", Price: 75},
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	sortBy := r.URL.Query().Get("sort_by") // "price" or "name"
	order := r.URL.Query().Get("order")     // "asc" or "desc"

	// Clone original slice to keep memory thread-safe
	result := make([]Product, len(products))
	copy(result, products)

	sort.Slice(result, func(i, j int) bool {
		if sortBy == "price" {
			if strings.ToLower(order) == "desc" {
				return result[i].Price > result[j].Price
			}
			return result[i].Price < result[j].Price
		}
		
		// Default sort by Name asc
		if strings.ToLower(order) == "desc" {
			return result[i].Name > result[j].Name
		}
		return result[i].Name < result[j].Name
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /products", sortHandler)
	http.ListenAndServe(":8080", mux)
}