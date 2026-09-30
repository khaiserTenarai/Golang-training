package repository
import (
	"context"
	"errors"
	"fmt"

	"employee-management/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrEmployeeNotFound = errors.New("employee not found")

type EmployeeRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(
	db *pgxpool.Pool,
) EmployeeRepository {

	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) Save(
	employee model.Employee,
) error {

	query := `
		INSERT INTO employees
		(name, age, email, salary, department_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.Salary,
		employee.Department.ID,
	)

	return err
}

func (r *EmployeeRepositoryImpl) FindByID(
	id int,
) (model.Employee, error) {

	var employee model.Employee

	query := `
		SELECT
			e.id,
			e.name,
			e.age,
			e.email,
			e.salary,
			d.id,
			d.name
		FROM employees e
		JOIN departments d
			ON e.department_id = d.id
		WHERE e.id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Age,
		&employee.Email,
		&employee.Salary,
		&employee.Department.ID,
		&employee.Department.Name,
	)

	if err == pgx.ErrNoRows {
		return employee, ErrEmployeeNotFound
	}

	return employee, err
}

func (r *EmployeeRepositoryImpl) FindAll() (
	[]model.Employee,
	error,
) {

	query := `
		SELECT
			e.id,
			e.name,
			e.age,
			e.email,
			e.salary,
			d.id,
			d.name
		FROM employees e
		JOIN departments d
			ON e.department_id = d.id
		ORDER BY e.id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employees := []model.Employee{}

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Age,
			&employee.Email,
			&employee.Salary,
			&employee.Department.ID,
			&employee.Department.Name,
		)

		if err != nil {
			return nil, err
		}

		employees = append(
			employees,
			employee,
		)
	}

	return employees, rows.Err()
}

func (r *EmployeeRepositoryImpl) Update(
	employee model.Employee,
) error {

	query := `
		UPDATE employees
		SET
			name = $1,
			age = $2,
			email = $3,
			salary = $4,
			department_id = $5
		WHERE id = $6
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Age,
		employee.Email,
		employee.Salary,
		employee.Department.ID,
		employee.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}

	return nil
}

func (r *EmployeeRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrEmployeeNotFound
	}

	return nil
}

func (r *EmployeeRepositoryImpl) Search(
	search model.EmployeeSearch,
) ([]model.Employee, int, error) {

	where := " WHERE 1=1 "
	args := []interface{}{}
	n := 1

	// Name
	if search.Name != "" {

		where += fmt.Sprintf(
			" AND e.name ILIKE $%d",
			n,
		)

		args = append(
			args,
			"%"+search.Name+"%",
		)

		n++
	}

	// Department
	if search.DepartmentID > 0 {

		where += fmt.Sprintf(
			" AND e.department_id = $%d",
			n,
		)

		args = append(
			args,
			search.DepartmentID,
		)

		n++
	}

	// Minimum salary
	if search.MinSalary > 0 {

		where += fmt.Sprintf(
			" AND e.salary >= $%d",
			n,
		)

		args = append(
			args,
			search.MinSalary,
		)

		n++
	}

	// Maximum salary
	if search.MaxSalary > 0 {

		where += fmt.Sprintf(
			" AND e.salary <= $%d",
			n,
		)

		args = append(
			args,
			search.MaxSalary,
		)

		n++
	}

	// Count
	countQuery := `
		SELECT COUNT(*)
		FROM employees e
	` + where

	var total int

	err := r.db.QueryRow(
		context.Background(),
		countQuery,
		args...,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	// Sorting
	sortBy := "e.id"

	if search.SortBy == "name" {
		sortBy = "e.name"
	}

	if search.SortBy == "salary" {
		sortBy = "e.salary"
	}

	if search.SortBy == "department" {
		sortBy = "d.name"
	}

	sortOrder := "ASC"

	if search.SortOrder == "desc" {
		sortOrder = "DESC"
	}

	query := `
		SELECT
			e.id,
			e.name,
			e.age,
			e.email,
			e.salary,
			d.id,
			d.name
		FROM employees e
		JOIN departments d
			ON e.department_id = d.id
	` + where

	query += " ORDER BY " + sortBy + " " + sortOrder

	offset := (search.Page - 1) * search.Size

	query += fmt.Sprintf(
		" LIMIT $%d OFFSET $%d",
		n,
		n+1,
	)

	args = append(
		args,
		search.Size,
		offset,
	)

	rows, err := r.db.Query(
		context.Background(),
		query,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	employees := []model.Employee{}

	for rows.Next() {

		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Age,
			&employee.Email,
			&employee.Salary,
			&employee.Department.ID,
			&employee.Department.Name,
		)

		if err != nil {
			return nil, 0, err
		}

		employees = append(
			employees,
			employee,
		)
	}

	return employees, total, rows.Err()
}
