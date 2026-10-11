package repository

import (
	"api/models"
	"database/sql"
	"fmt"
	"log"
)

type EmployeeRepository interface {
	GetEmployees(page, limit int, deptFilter, sortBy string) []models.Employee
	Save(emp models.Employee) models.Employee
}

type employeeRepoImpl struct {
	db *sql.DB
}

// Now requires a database connection to be injected
func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepoImpl{db: db}
}

func (r *employeeRepoImpl) GetEmployees(page, limit int, deptFilter, sortBy string) []models.Employee {
	var employees []models.Employee
	
	// Base query
	query := "SELECT id, name, department, role FROM employees"
	var args []interface{}
	argId := 1

	// 11. Filtering
	if deptFilter != "" {
		query += fmt.Sprintf(" WHERE department = $%d", argId)
		args = append(args, deptFilter)
		argId++
	}

	// 12. Sorting (preventing SQL injection by strictly checking the input)
	if sortBy == "name" {
		query += " ORDER BY name ASC"
	} else {
		query += " ORDER BY id ASC" // default sort
	}

	// 10. Pagination
	offset := (page - 1) * limit
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argId, argId+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return employees
	}
	defer rows.Close()

	for rows.Next() {
		var e models.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Department, &e.Role); err == nil {
			employees = append(employees, e)
		}
	}
	return employees
}

func (r *employeeRepoImpl) Save(emp models.Employee) models.Employee {
	query := `INSERT INTO employees (name, department, role) VALUES ($1, $2, $3) RETURNING id`
	
	err := r.db.QueryRow(query, emp.Name, emp.Department, emp.Role).Scan(&emp.ID)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
	}
	return emp
}