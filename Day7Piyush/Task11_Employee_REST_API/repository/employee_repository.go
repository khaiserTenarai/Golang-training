package repository

import (
	"database/sql"
	"task11_employee_rest_api/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(emp models.Employee) (models.Employee, error) {
	query := `INSERT INTO rest_employees (name, email, age, department, salary) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, age, department, salary, created_at`
	err := r.DB.QueryRow(query, emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *EmployeeRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, age, department, salary, created_at FROM rest_employees ORDER BY id`)
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
	err := r.DB.QueryRow(`SELECT id, name, email, age, department, salary, created_at FROM rest_employees WHERE id=$1`, id).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *EmployeeRepository) Update(emp models.Employee) (models.Employee, error) {
	query := `UPDATE rest_employees SET name=$1, email=$2, age=$3, department=$4, salary=$5 WHERE id=$6
		RETURNING id, name, email, age, department, salary, created_at`
	err := r.DB.QueryRow(query, emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary, emp.ID).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

func (r *EmployeeRepository) Delete(id int) error {
	result, err := r.DB.Exec(`DELETE FROM rest_employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
