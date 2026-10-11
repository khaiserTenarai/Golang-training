package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"task14_15_employee_api/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`INSERT INTO employees (name, email, age, department, salary)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, age, department, salary, created_at, updated_at`,
		emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department,
			&emp.Salary, &emp.CreatedAt, &emp.UpdatedAt)
	return emp, err
}

func (r *EmployeeRepository) GetAll(name, department string, minSalary, maxSalary float64, page, pageSize int) ([]models.Employee, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	where := []string{}
	args := []interface{}{}
	idx := 1

	if name != "" {
		where = append(where, fmt.Sprintf("name ILIKE $%d", idx))
		args = append(args, "%"+name+"%")
		idx++
	}
	if department != "" {
		where = append(where, fmt.Sprintf("department ILIKE $%d", idx))
		args = append(args, department)
		idx++
	}
	if minSalary > 0 {
		where = append(where, fmt.Sprintf("salary >= $%d", idx))
		args = append(args, minSalary)
		idx++
	}
	if maxSalary > 0 {
		where = append(where, fmt.Sprintf("salary <= $%d", idx))
		args = append(args, maxSalary)
		idx++
	}

	whereClause := ""
	if len(where) > 0 {
		whereClause = "WHERE " + strings.Join(where, " AND ")
	}

	var total int
	r.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM employees %s", whereClause), args...).Scan(&total)

	offset := (page - 1) * pageSize
	query := fmt.Sprintf(
		"SELECT id, name, email, age, department, salary, created_at, updated_at FROM employees %s ORDER BY id LIMIT $%d OFFSET $%d",
		whereClause, idx, idx+1)
	args = append(args, pageSize, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department,
			&emp.Salary, &emp.CreatedAt, &emp.UpdatedAt); err != nil {
			return nil, 0, err
		}
		employees = append(employees, emp)
	}
	return employees, total, rows.Err()
}

func (r *EmployeeRepository) GetByID(id int) (models.Employee, error) {
	var emp models.Employee
	err := r.DB.QueryRow(
		`SELECT id, name, email, age, department, salary, created_at, updated_at
		FROM employees WHERE id = $1`, id).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department,
			&emp.Salary, &emp.CreatedAt, &emp.UpdatedAt)
	return emp, err
}

func (r *EmployeeRepository) Update(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`UPDATE employees SET name=$1, email=$2, age=$3, department=$4, salary=$5, updated_at=NOW()
		WHERE id=$6
		RETURNING id, name, email, age, department, salary, created_at, updated_at`,
		emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary, emp.ID).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department,
			&emp.Salary, &emp.CreatedAt, &emp.UpdatedAt)
	return emp, err
}

func (r *EmployeeRepository) Delete(id int) error {
	result, err := r.DB.Exec(`DELETE FROM employees WHERE id = $1`, id)
	if err != nil {
		return err
	}
	n, _ := result.RowsAffected()
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
