package repository

import (
	"database/sql"
	"employee_management/config"
	"employee_management/model"
)

type EmployeeRepository struct{}

// Let's filter by department now instead of position
func (r *EmployeeRepository) FindAll(department string) ([]model.Employee, error) {
	var employees []model.Employee
	var rows *sql.Rows
	var err error

	if department != "" {
		query := "SELECT id, name, department, role FROM employees WHERE department = $1"
		rows, err = config.DB.Query(query, department)
	} else {
		query := "SELECT id, name, department, role FROM employees"
		rows, err = config.DB.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Role); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func (r *EmployeeRepository) FindByID(id string) (model.Employee, error) {
	var emp model.Employee
	query := "SELECT id, name, department, role FROM employees WHERE id = $1"
	err := config.DB.QueryRow(query, id).Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Role)
	return emp, err
}

func (r *EmployeeRepository) Create(employee *model.Employee) error {
	query := "INSERT INTO employees (name, department, role) VALUES ($1, $2, $3) RETURNING id"
	return config.DB.QueryRow(query, employee.Name, employee.Department, employee.Role).Scan(&employee.ID)
}

func (r *EmployeeRepository) Update(employee *model.Employee) error {
	query := "UPDATE employees SET name = $1, department = $2, role = $3 WHERE id = $4"
	_, err := config.DB.Exec(query, employee.Name, employee.Department, employee.Role, employee.ID)
	return err
}

func (r *EmployeeRepository) Delete(id string) error {
	query := "DELETE FROM employees WHERE id = $1"
	_, err := config.DB.Exec(query, id)
	return err
}