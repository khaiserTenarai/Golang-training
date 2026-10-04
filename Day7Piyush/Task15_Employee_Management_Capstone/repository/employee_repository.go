package repository

import (
	"database/sql"
	"fmt"
	"math"
	"strings"
	"task15_employee_management_capstone/models"
)

// EmployeeRepository interface
type EmployeeRepository interface {
	Create(emp models.Employee) (models.Employee, error)
	GetAll(page, pageSize int, search, sortBy, sortOrder string) ([]models.Employee, int, error)
	GetByID(id int) (models.Employee, error)
	Update(emp models.Employee) (models.Employee, error)
	Delete(id int) error
	UpdateSalary(empID int, newSalary float64, reason string) error
	GetSalaryHistory(empID int) ([]models.SalaryHistory, error)
}

// PostgresEmployeeRepository implements EmployeeRepository
type PostgresEmployeeRepository struct {
	DB *sql.DB
}

func NewPostgresEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &PostgresEmployeeRepository{DB: db}
}

func (r *PostgresEmployeeRepository) Create(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`INSERT INTO cap_employees (name, email, age, department_id, salary)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, age, department_id, salary, created_at, updated_at`,
		emp.Name, emp.Email, emp.Age, emp.DepartmentID, emp.Salary).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.DepartmentID, &emp.Salary, &emp.CreatedAt, &emp.UpdatedAt)
	return emp, err
}

func (r *PostgresEmployeeRepository) GetAll(page, pageSize int, search, sortBy, sortOrder string) ([]models.Employee, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// Build WHERE clause
	where := []string{}
	args := []interface{}{}
	idx := 1
	if search != "" {
		where = append(where, fmt.Sprintf("(e.name ILIKE $%d OR e.email ILIKE $%d)", idx, idx))
		args = append(args, "%"+search+"%")
		idx++
	}
	whereClause := ""
	if len(where) > 0 {
		whereClause = "WHERE " + strings.Join(where, " AND ")
	}

	// Count total
	var total int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM cap_employees e %s", whereClause)
	r.DB.QueryRow(countQuery, args...).Scan(&total)

	// Sort
	validSort := map[string]bool{"name": true, "email": true, "salary": true, "id": true, "created_at": true}
	if !validSort[sortBy] {
		sortBy = "e.id"
	} else {
		sortBy = "e." + sortBy
	}
	if strings.ToUpper(sortOrder) != "DESC" {
		sortOrder = "ASC"
	}

	// Pagination
	offset := (page - 1) * pageSize
	query := fmt.Sprintf(
		`SELECT e.id, e.name, e.email, e.age, e.department_id, COALESCE(d.name, 'Unassigned'),
		e.salary, e.created_at, e.updated_at
		FROM cap_employees e
		LEFT JOIN cap_departments d ON e.department_id = d.id
		%s ORDER BY %s %s LIMIT $%d OFFSET $%d`,
		whereClause, sortBy, sortOrder, idx, idx+1)
	args = append(args, pageSize, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.DepartmentID,
			&emp.DepartmentName, &emp.Salary, &emp.CreatedAt, &emp.UpdatedAt); err != nil {
			return nil, 0, err
		}
		employees = append(employees, emp)
	}
	return employees, total, rows.Err()
}

func (r *PostgresEmployeeRepository) GetByID(id int) (models.Employee, error) {
	var emp models.Employee
	err := r.DB.QueryRow(
		`SELECT e.id, e.name, e.email, e.age, e.department_id, COALESCE(d.name, 'Unassigned'),
		e.salary, e.created_at, e.updated_at
		FROM cap_employees e
		LEFT JOIN cap_departments d ON e.department_id = d.id
		WHERE e.id = $1`, id).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.DepartmentID,
			&emp.DepartmentName, &emp.Salary, &emp.CreatedAt, &emp.UpdatedAt)
	return emp, err
}

func (r *PostgresEmployeeRepository) Update(emp models.Employee) (models.Employee, error) {
	err := r.DB.QueryRow(
		`UPDATE cap_employees SET name=$1, email=$2, age=$3, department_id=$4, salary=$5, updated_at=NOW()
		WHERE id=$6
		RETURNING id, name, email, age, department_id, salary, created_at, updated_at`,
		emp.Name, emp.Email, emp.Age, emp.DepartmentID, emp.Salary, emp.ID).
		Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.DepartmentID, &emp.Salary, &emp.CreatedAt, &emp.UpdatedAt)
	return emp, err
}

func (r *PostgresEmployeeRepository) Delete(id int) error {
	result, err := r.DB.Exec(`DELETE FROM cap_employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	n, _ := result.RowsAffected()
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// UpdateSalary uses a transaction for atomic salary update + history record
func (r *PostgresEmployeeRepository) UpdateSalary(empID int, newSalary float64, reason string) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	var oldSalary float64
	err = tx.QueryRow(`SELECT salary FROM cap_employees WHERE id=$1 FOR UPDATE`, empID).Scan(&oldSalary)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("employee not found: %w", err)
	}

	_, err = tx.Exec(`UPDATE cap_employees SET salary=$1, updated_at=NOW() WHERE id=$2`, newSalary, empID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("update salary: %w", err)
	}

	_, err = tx.Exec(
		`INSERT INTO cap_salary_history (employee_id, old_salary, new_salary, reason) VALUES ($1, $2, $3, $4)`,
		empID, oldSalary, newSalary, reason)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("record history: %w", err)
	}

	return tx.Commit()
}

func (r *PostgresEmployeeRepository) GetSalaryHistory(empID int) ([]models.SalaryHistory, error) {
	rows, err := r.DB.Query(
		`SELECT id, employee_id, old_salary, new_salary, reason, changed_at FROM cap_salary_history
		WHERE employee_id=$1 ORDER BY changed_at DESC`, empID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var history []models.SalaryHistory
	for rows.Next() {
		var h models.SalaryHistory
		if err := rows.Scan(&h.ID, &h.EmployeeID, &h.OldSalary, &h.NewSalary, &h.Reason, &h.ChangedAt); err != nil {
			return nil, err
		}
		history = append(history, h)
	}
	return history, rows.Err()
}

// Helper
func TotalPages(total, pageSize int) int {
	return int(math.Ceil(float64(total) / float64(pageSize)))
}
