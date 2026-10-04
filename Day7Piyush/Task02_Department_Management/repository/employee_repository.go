package repository

import (
	"database/sql"
	"task02_department_management/models"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(emp models.Employee) (int, error) {
	var id int
	query := `INSERT INTO dept_employees (name, email, department_id, salary) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, emp.Name, emp.Email, emp.DepartmentID, emp.Salary).Scan(&id)
	return id, err
}

func (r *EmployeeRepository) GetAllWithDepartment() ([]models.Employee, error) {
	query := `SELECT e.id, e.name, e.email, e.department_id, COALESCE(d.name, 'Unassigned'), e.salary, e.created_at
		FROM dept_employees e
		LEFT JOIN departments d ON e.department_id = d.id
		ORDER BY e.id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.DepartmentID, &emp.DepartmentName, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, rows.Err()
}

func (r *EmployeeRepository) GetByDepartment(deptID int) ([]models.Employee, error) {
	query := `SELECT e.id, e.name, e.email, e.department_id, d.name, e.salary, e.created_at
		FROM dept_employees e
		JOIN departments d ON e.department_id = d.id
		WHERE e.department_id = $1 ORDER BY e.id`
	rows, err := r.DB.Query(query, deptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.DepartmentID, &emp.DepartmentName, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, rows.Err()
}

func (r *EmployeeRepository) AssignDepartment(empID, deptID int) error {
	_, err := r.DB.Exec(`UPDATE dept_employees SET department_id=$1 WHERE id=$2`, deptID, empID)
	return err
}

func (r *EmployeeRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM dept_employees WHERE id=$1`, id)
	return err
}
