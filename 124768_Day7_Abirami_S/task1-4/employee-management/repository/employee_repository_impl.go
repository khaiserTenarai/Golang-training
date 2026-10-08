package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"employee-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) CreateEmployee(employee model.Employee) error {

	fmt.Println("\n----- CREATE EMPLOYEE -----")

	query := `
		INSERT INTO employees (name, email, age, salary, department_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.DepartmentID,
	)

	if err != nil {
		return err
	}

	fmt.Println("Employee created successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *EmployeeRepositoryImpl) GetEmployee(id int) (*model.Employee, error) {

	fmt.Println("\n----- READ EMPLOYEE -----")

	var employee model.Employee

	query := `
		SELECT id, name, email, age, salary, COALESCE(department_id,0)
		FROM employees
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Age,
		&employee.Salary,
		&employee.DepartmentID,
	)

	if err != nil {

		if err == pgx.ErrNoRows {
			return nil, ErrEmployeeNotFound
		}

		return nil, err
	}

	fmt.Println("ID           :", employee.ID)
	fmt.Println("Name         :", employee.Name)
	fmt.Println("Email        :", employee.Email)
	fmt.Println("Age          :", employee.Age)
	fmt.Println("Salary       :", employee.Salary)
	fmt.Println("Department ID:", employee.DepartmentID)

	return &employee, nil
}

func (r *EmployeeRepositoryImpl) GetAllEmployees() ([]model.Employee, error) {

	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	query := `
		SELECT id, name, email, age, salary, COALESCE(department_id,0)
		FROM employees
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Age,
			&employee.Salary,
			&employee.DepartmentID,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepositoryImpl) UpdateEmployee(employee model.Employee) error {

	fmt.Println("\n----- UPDATE EMPLOYEE -----")

	query := `
		UPDATE employees
		SET name = $1,
		    email = $2,
		    age = $3,
		    salary = $4,
		    department_id = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.DepartmentID,
		employee.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}

	fmt.Println("Employee updated successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *EmployeeRepositoryImpl) DeleteEmployee(id int) error {

	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err := r.db.Exec(
		context.Background(),
		`DELETE FROM employees WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}

	fmt.Println("Employee deleted successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

func (r *EmployeeRepositoryImpl) SearchEmployees(name string, departmentID int, salary float64, page int, size int, sortBy string, sortOrder string) ([]model.Employee, error) {

	fmt.Println("\n----- SEARCH EMPLOYEES -----")

	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 5
	}

	offset := (page - 1) * size

	sortColumns := map[string]string{
		"id":     "e.id",
		"name":   "e.name",
		"age":    "e.age",
		"salary": "e.salary",
	}

	sortColumn, exists := sortColumns[strings.ToLower(sortBy)]

	if !exists {
		sortColumn = "e.id"
	}

	sortOrder = strings.ToUpper(sortOrder)

	if sortOrder != "ASC" && sortOrder != "DESC" {
		sortOrder = "ASC"
	}

	query := `
		SELECT e.id, e.name, e.email, e.age, e.salary, COALESCE(e.department_id,0)
		FROM employees e
		LEFT JOIN departments d ON e.department_id = d.id
		WHERE 1 = 1
	`

	var args []interface{}
	argNumber := 1

	if strings.TrimSpace(name) != "" {
		query += fmt.Sprintf(" AND e.name ILIKE $%d", argNumber)
		args = append(args, "%"+name+"%")
		argNumber++
	}

	if departmentID > 0 {
		query += fmt.Sprintf(" AND d.id = $%d", argNumber)
		args = append(args, departmentID)
		argNumber++
	}

	if salary > 0 {
		query += fmt.Sprintf(" AND e.salary >= $%d", argNumber)
		args = append(args, salary)
		argNumber++
	}

	query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", sortColumn, sortOrder, argNumber, argNumber+1)

	args = append(args, size, offset)

	rows, err := r.db.Query(
		context.Background(),
		query,
		args...,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Age,
			&employee.Salary,
			&employee.DepartmentID,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
