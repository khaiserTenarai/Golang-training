package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type EmployeeSearchResult struct {
	ID             int
	Name           string
	Position       string
	Salary         float64
	DepartmentName string
}

type SearchFilter struct {
	Name       string
	Department string
	MinSalary  float64
	SortBy     string
	SortOrder  string
	Page       int
	PageSize   int
}

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pgadmin dbname=gotraining sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	filter := SearchFilter{
		Name:       "a",
		Department: "Engineering",
		MinSalary:  80000,
		SortBy:     "salary",
		SortOrder:  "DESC",
		Page:       1,
		PageSize:   5,
	}

	results, err := searchEmployees(db, filter)
	if err != nil {
		log.Fatalf("Search failed: %v", err)
	}

	fmt.Printf("Found %d employees:\n", len(results))
	for _, emp := range results {
		fmt.Printf("- %s (Dept: %s) | %s | $%.2f\n", emp.Name, emp.DepartmentName, emp.Position, emp.Salary)
	}
}

func searchEmployees(db *sql.DB, filter SearchFilter) ([]EmployeeSearchResult, error) {
	query := `
		SELECT e.id, e.name, e.position, e.salary, d.name as department_name 
		FROM employees e
		JOIN departments d ON e.department_id = d.id
		WHERE 1=1
	`
	args := []interface{}{}
	argId := 1

	if filter.Name != "" {
		query += fmt.Sprintf(" AND e.name ILIKE $%d", argId)
		args = append(args, "%"+filter.Name+"%")
		argId++
	}
	if filter.Department != "" {
		query += fmt.Sprintf(" AND d.name ILIKE $%d", argId)
		args = append(args, "%"+filter.Department+"%")
		argId++
	}
	if filter.MinSalary > 0 {
		query += fmt.Sprintf(" AND e.salary >= $%d", argId)
		args = append(args, filter.MinSalary)
		argId++
	}

	sortBy := "e.id"
	switch strings.ToLower(filter.SortBy) {
	case "name":
		sortBy = "e.name"
	case "salary":
		sortBy = "e.salary"
	case "department":
		sortBy = "d.name"
	}

	sortOrder := "ASC"
	if strings.ToUpper(filter.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	query += fmt.Sprintf(" ORDER BY %s %s", sortBy, sortOrder)

	pageSize := filter.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	page := filter.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, pageSize, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []EmployeeSearchResult
	for rows.Next() {
		var emp EmployeeSearchResult
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary, &emp.DepartmentName); err != nil {
			return nil, err
		}
		results = append(results, emp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}