package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"task03_employee_search/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Insert(emp models.Employee) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO search_employees (name, department, salary) VALUES ($1, $2, $3) RETURNING id`,
		emp.Name, emp.Department, emp.Salary).Scan(&id)
	return id, err
}

func (r *EmployeeRepository) Search(params models.SearchParams) ([]models.Employee, int, error) {
	where := []string{}
	args := []interface{}{}
	idx := 1

	if params.Name != "" {
		where = append(where, fmt.Sprintf("name ILIKE $%d", idx))
		args = append(args, "%"+params.Name+"%")
		idx++
	}
	if params.Department != "" {
		where = append(where, fmt.Sprintf("department ILIKE $%d", idx))
		args = append(args, "%"+params.Department+"%")
		idx++
	}
	if params.MinSalary > 0 {
		where = append(where, fmt.Sprintf("salary >= $%d", idx))
		args = append(args, params.MinSalary)
		idx++
	}
	if params.MaxSalary > 0 {
		where = append(where, fmt.Sprintf("salary <= $%d", idx))
		args = append(args, params.MaxSalary)
		idx++
	}

	whereClause := ""
	if len(where) > 0 {
		whereClause = "WHERE " + strings.Join(where, " AND ")
	}

	// Count total results
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM search_employees %s", whereClause)
	var total int
	if err := r.DB.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Sorting
	validSort := map[string]bool{"name": true, "department": true, "salary": true, "id": true}
	sortBy := "id"
	if validSort[params.SortBy] {
		sortBy = params.SortBy
	}
	sortOrder := "ASC"
	if strings.ToUpper(params.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	// Pagination
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 {
		params.PageSize = 5
	}
	offset := (params.Page - 1) * params.PageSize

	query := fmt.Sprintf(
		"SELECT id, name, department, salary, created_at FROM search_employees %s ORDER BY %s %s LIMIT $%d OFFSET $%d",
		whereClause, sortBy, sortOrder, idx, idx+1)
	args = append(args, params.PageSize, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, 0, err
		}
		employees = append(employees, emp)
	}
	return employees, total, rows.Err()
}
