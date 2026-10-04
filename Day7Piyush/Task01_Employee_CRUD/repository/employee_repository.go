package repository

import (
	"database/sql"
	"task01_employee_crud/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(emp models.Employee) (int, error) {
	var id int
	query := `INSERT INTO employees (name, email, age, department, salary) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary).Scan(&id)
	return id, err
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	query := `SELECT id, name, email, age, department, salary, created_at FROM employees ORDER BY id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, rows.Err()
}

func (r *EmployeeRepository) GetByID(id int) (models.Employee, error) {
	var emp models.Employee
	query := `SELECT id, name, email, age, department, salary, created_at FROM employees WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *EmployeeRepository) Update(emp models.Employee) error {
	query := `UPDATE employees SET name=$1, email=$2, age=$3, department=$4, salary=$5 WHERE id=$6`
	_, err := r.DB.Exec(query, emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary, emp.ID)
	return err
}

func (r *EmployeeRepository) Delete(id int) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
