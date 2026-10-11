package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type ListResponse struct {
	Data []Employee `json:"data"`
	Meta ListMeta   `json:"meta"`
}

type ListMeta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

var employees = []Employee{
	{ID: 1, Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000},
	{ID: 2, Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000},
	{ID: 3, Name: "Sara Khan", Email: "sara@example.com", Department: "Engineering", Salary: 91000},
	{ID: 4, Name: "Vikram Singh", Email: "vikram@example.com", Department: "HR", Salary: 58000},
	{ID: 5, Name: "Meera Nair", Email: "meera@example.com", Department: "Sales", Salary: 54000},
	{ID: 6, Name: "Arjun Mehta", Email: "arjun@example.com", Department: "Engineering", Salary: 76000},
	{ID: 7, Name: "Priya Sharma", Email: "priya@example.com", Department: "Finance", Salary: 69000},
	{ID: 8, Name: "Karan Patel", Email: "karan@example.com", Department: "Sales", Salary: 51000},
	{ID: 9, Name: "Divya Iyer", Email: "divya@example.com", Department: "Marketing", Salary: 63000},
	{ID: 10, Name: "Nikhil Das", Email: "nikhil@example.com", Department: "Engineering", Salary: 99000},
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func intParam(q map[string][]string, key string, def int) (int, error) {
	values, ok := q[key]
	if !ok || values[0] == "" {
		return def, nil
	}
	n, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", key)
	}
	return n, nil
}

func floatParam(q map[string][]string, key string) (float64, bool, error) {
	values, ok := q[key]
	if !ok || values[0] == "" {
		return 0, false, nil
	}
	n, err := strconv.ParseFloat(values[0], 64)
	if err != nil {
		return 0, false, fmt.Errorf("%s must be a number", key)
	}
	return n, true, nil
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	q := r.URL.Query()

	page, err := intParam(q, "page", 1)
	if err != nil || page < 1 {
		writeError(w, http.StatusBadRequest, "page must be a positive integer")
		return
	}

	limit, err := intParam(q, "limit", 5)
	if err != nil || limit < 1 || limit > 50 {
		writeError(w, http.StatusBadRequest, "limit must be between 1 and 50")
		return
	}

	minSalary, hasMin, err := floatParam(q, "min_salary")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	maxSalary, hasMax, err := floatParam(q, "max_salary")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	sortBy := q.Get("sort")
	if sortBy == "" {
		sortBy = "id"
	}
	order := strings.ToLower(q.Get("order"))
	if order == "" {
		order = "asc"
	}
	if order != "asc" && order != "desc" {
		writeError(w, http.StatusBadRequest, "order must be asc or desc")
		return
	}

	less := map[string]func(a, b Employee) bool{
		"id":         func(a, b Employee) bool { return a.ID < b.ID },
		"name":       func(a, b Employee) bool { return a.Name < b.Name },
		"salary":     func(a, b Employee) bool { return a.Salary < b.Salary },
		"department": func(a, b Employee) bool { return a.Department < b.Department },
	}[sortBy]
	if less == nil {
		writeError(w, http.StatusBadRequest, "sort must be one of: id, name, salary, department")
		return
	}

	department := q.Get("department")
	name := strings.ToLower(q.Get("name"))

	filtered := make([]Employee, 0, len(employees))
	for _, e := range employees {
		if department != "" && !strings.EqualFold(e.Department, department) {
			continue
		}
		if name != "" && !strings.Contains(strings.ToLower(e.Name), name) {
			continue
		}
		if hasMin && e.Salary < minSalary {
			continue
		}
		if hasMax && e.Salary > maxSalary {
			continue
		}
		filtered = append(filtered, e)
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		if order == "desc" {
			return less(filtered[j], filtered[i])
		}
		return less(filtered[i], filtered[j])
	})

	total := len(filtered)
	totalPages := (total + limit - 1) / limit
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}

	writeJSON(w, http.StatusOK, ListResponse{
		Data: filtered[start:end],
		Meta: ListMeta{Page: page, Limit: limit, Total: total, TotalPages: totalPages},
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employeesHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
