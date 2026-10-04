package repository

import (
	"database/sql"
	"fmt"
	"task04_salary_management/models"
)

type SalaryRepository struct {
	DB *sql.DB
}

func NewSalaryRepository(db *sql.DB) *SalaryRepository {
	return &SalaryRepository{DB: db}
}

func (r *SalaryRepository) CreateEmployee(emp models.Employee) (int, error) {
	var id int
	err := r.DB.QueryRow(
		`INSERT INTO sal_employees (name, department, salary) VALUES ($1, $2, $3) RETURNING id`,
		emp.Name, emp.Department, emp.Salary).Scan(&id)
	return id, err
}

func (r *SalaryRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.DB.Query(`SELECT id, name, department, salary, created_at FROM sal_employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Salary, &emp.CreatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, rows.Err()
}

func (r *SalaryRepository) GetByID(id int) (models.Employee, error) {
	var emp models.Employee
	err := r.DB.QueryRow(`SELECT id, name, department, salary, created_at FROM sal_employees WHERE id=$1`, id).
		Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Salary, &emp.CreatedAt)
	return emp, err
}

// UpdateSalary uses a PostgreSQL transaction to update salary and record history
func (r *SalaryRepository) UpdateSalary(empID int, newSalary float64, reason string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	// Get current salary within the transaction
	var oldSalary float64
	err = tx.QueryRow(`SELECT salary FROM sal_employees WHERE id = $1 FOR UPDATE`, empID).Scan(&oldSalary)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("employee not found: %w", err)
	}

	// Update the employee's salary
	_, err = tx.Exec(`UPDATE sal_employees SET salary = $1 WHERE id = $2`, newSalary, empID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update salary: %w", err)
	}

	// Record in salary history
	_, err = tx.Exec(
		`INSERT INTO salary_history (employee_id, old_salary, new_salary, change_reason) VALUES ($1, $2, $3, $4)`,
		empID, oldSalary, newSalary, reason)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("record history: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}

func (r *SalaryRepository) GetSalaryHistory(empID int) ([]models.SalaryHistory, error) {
	query := `SELECT sh.id, sh.employee_id, e.name, sh.old_salary, sh.new_salary, sh.change_reason, sh.changed_at
		FROM salary_history sh
		JOIN sal_employees e ON sh.employee_id = e.id
		WHERE sh.employee_id = $1
		ORDER BY sh.changed_at DESC`
	rows, err := r.DB.Query(query, empID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.SalaryHistory
	for rows.Next() {
		var h models.SalaryHistory
		if err := rows.Scan(&h.ID, &h.EmployeeID, &h.EmployeeName, &h.OldSalary, &h.NewSalary, &h.ChangeReason, &h.ChangedAt); err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	return history, rows.Err()
}
