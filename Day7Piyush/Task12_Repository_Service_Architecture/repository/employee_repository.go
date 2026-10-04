package repository

import (
	"database/sql"
	"task12_repository_service_architecture/models"
)

// EmployeeRepository interface - defines the contract for data access
type EmployeeRepository interface {
	Create(emp models.Employee) (models.Employee, error)
	GetAll() ([]models.Employee, error)
	GetByID(id int) (models.Employee, error)
	Update(emp models.Employee) (models.Employee, error)
	Delete(id int) error
}

// PostgresEmployeeRepository implements EmployeeRepository using PostgreSQL
type PostgresEmployeeRepository struct {
	DB *sql.DB
}

func NewPostgresEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &PostgresEmployeeRepository{DB: db}
}

func (r *PostgresEmployeeRepository) Create(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`INSERT INTO svc_employees (name, email, department, salary) VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, department, salary, created_at`,
		emp.Name, emp.Email, emp.Department, emp.Salary).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *PostgresEmployeeRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, department, salary, created_at FROM svc_employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, rows.Err()
}

func (r *PostgresEmployeeRepository) GetByID(id int) (models.Employee, error) {
	var emp models.Employee
	err := r.DB.QueryRow(
		`SELECT id, name, email, department, salary, created_at FROM svc_employees WHERE id=$1`, id).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *PostgresEmployeeRepository) Update(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`UPDATE svc_employees SET name=$1, email=$2, department=$3, salary=$4 WHERE id=$5
		RETURNING id, name, email, department, salary, created_at`,
		emp.Name, emp.Email, emp.Department, emp.Salary, emp.ID).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *PostgresEmployeeRepository) Delete(id int) error {
	result, err := r.DB.Exec(`DELETE FROM svc_employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
