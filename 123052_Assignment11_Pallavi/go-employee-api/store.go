package main

import (
	"database/sql"
)

// Store interface defines all database operations
type Store interface {
	GetAllEmployees(role string) ([]Employee, error)
	GetEmployeeByID(id int) (*Employee, error)
	CreateEmployee(emp *Employee) error
	UpdateEmployee(id int, emp *Employee) error
	DeleteEmployee(id int) error
}

// PostgresStore implements the Store interface using PostgreSQL
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore initializes a new Store instance
func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

// GetAllEmployees fetches all employees, optionally filtered by role
func (s *PostgresStore) GetAllEmployees(role string) ([]Employee, error) {
	var rows *sql.Rows
	var err error

	if role != "" {
		rows, err = s.db.Query("SELECT id, name, role FROM employees WHERE role = $1", role)
	} else {
		rows, err = s.db.Query("SELECT id, name, role FROM employees")
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var e Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Role); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

// GetEmployeeByID retrieves a single employee by primary key
func (s *PostgresStore) GetEmployeeByID(id int) (*Employee, error) {
	var e Employee
	err := s.db.QueryRow("SELECT id, name, role FROM employees WHERE id = $1", id).Scan(&e.ID, &e.Name, &e.Role)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// CreateEmployee inserts a new employee record and returns the assigned ID
func (s *PostgresStore) CreateEmployee(emp *Employee) error {
	query := "INSERT INTO employees (name, role) VALUES ($1, $2) RETURNING id"
	return s.db.QueryRow(query, emp.Name, emp.Role).Scan(&emp.ID)
}

// UpdateEmployee modifies an existing employee record
func (s *PostgresStore) UpdateEmployee(id int, emp *Employee) error {
	query := "UPDATE employees SET name = $1, role = $2 WHERE id = $3"
	res, err := s.db.Exec(query, emp.Name, emp.Role, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	emp.ID = id
	return nil
}

// DeleteEmployee removes an employee record by ID
func (s *PostgresStore) DeleteEmployee(id int) error {
	query := "DELETE FROM employees WHERE id = $1"
	res, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}