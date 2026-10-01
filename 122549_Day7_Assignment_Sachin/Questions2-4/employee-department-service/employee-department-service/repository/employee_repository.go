package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"employee-department-service/model"
)

// ErrEmployeeNotFound is returned when an employee ID doesn't exist.
var ErrEmployeeNotFound = errors.New("employee not found")

// CreateEmployee inserts a new employee row. emp.DepartmentID can be
// nil if the employee isn't assigned to a department yet.
func CreateEmployee(db *sql.DB, emp model.Employee) (int, error) {
	var id int
	query := `
		INSERT INTO new_employee (name, email, age, salary, department_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := db.QueryRow(query, emp.Name, emp.Email, emp.Age, emp.Salary, emp.DepartmentID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create employee failed: %w", err)
	}
	return id, nil
}

// GetEmployee fetches one employee by ID.
func GetEmployee(db *sql.DB, id int) (model.Employee, error) {
	var e model.Employee
	query := `
		SELECT id, name, email, age, salary, department_id, created_at
		FROM new_employee
		WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Salary, &e.DepartmentID, &e.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Employee{}, fmt.Errorf("get employee failed: %w", ErrEmployeeNotFound)
		}
		return model.Employee{}, fmt.Errorf("get employee failed: %w", err)
	}
	return e, nil
}

// DeleteEmployee removes an employee by ID.
func DeleteEmployee(db *sql.DB, id int) error {
	result, err := db.Exec(`DELETE FROM new_employee WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete employee failed: %w", err)
	}
	return checkRowsAffected(result, ErrEmployeeNotFound)
}

// EmployeeSearchParams holds every optional filter/sort/pagination
// option for SearchEmployees. Leave a field at its zero value to skip
// that filter - e.g. Name == "" means "don't filter by name".
type EmployeeSearchParams struct {
	Name           string  // partial, case-insensitive match on employee name
	DepartmentName string  // partial, case-insensitive match on department name
	MinSalary      float64 // 0 means "no minimum"
	MaxSalary      float64 // 0 means "no maximum"
	SortBy         string  // "name", "salary", "age" or "" (defaults to id)
	SortOrder      string  // "asc" or "desc" (defaults to "asc")
	Page           int     // 1-based; defaults to 1
	PageSize       int     // defaults to 10
}

// sortColumns whitelists which columns are allowed in ORDER BY.
// ORDER BY can't be parameterized like a normal value ($1, $2, ...),
// so we never build it directly from raw user input - only from this
// fixed map. That's what keeps this safe from SQL injection.
var sortColumns = map[string]string{
	"name":   "e.name",
	"salary": "e.salary",
	"age":    "e.age",
	"id":     "e.id",
}

// SearchEmployees returns one page of employees matching the given
// filters, plus the total number of matches (ignoring pagination) so
// the caller can work out how many pages exist.
func SearchEmployees(db *sql.DB, params EmployeeSearchParams) ([]model.Employee, int, error) {
	var conditions []string
	var args []interface{}
	argPos := 1

	baseFrom := `FROM new_employee e LEFT JOIN department d ON e.department_id = d.id`

	if params.Name != "" {
		conditions = append(conditions, fmt.Sprintf("e.name ILIKE $%d", argPos))
		args = append(args, "%"+params.Name+"%")
		argPos++
	}
	if params.DepartmentName != "" {
		conditions = append(conditions, fmt.Sprintf("d.name ILIKE $%d", argPos))
		args = append(args, "%"+params.DepartmentName+"%")
		argPos++
	}
	if params.MinSalary > 0 {
		conditions = append(conditions, fmt.Sprintf("e.salary >= $%d", argPos))
		args = append(args, params.MinSalary)
		argPos++
	}
	if params.MaxSalary > 0 {
		conditions = append(conditions, fmt.Sprintf("e.salary <= $%d", argPos))
		args = append(args, params.MaxSalary)
		argPos++
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	// Total count first, before pagination is applied - this is what
	// lets the caller know how many pages of results there are.
	countQuery := fmt.Sprintf("SELECT COUNT(*) %s %s", baseFrom, whereClause)
	var total int
	if err := db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count employees failed: %w", err)
	}

	// Sorting: look the requested column up in the whitelist above,
	// and fall back to sorting by ID if it's not a recognised value.
	sortColumn, ok := sortColumns[strings.ToLower(params.SortBy)]
	if !ok {
		sortColumn = "e.id"
	}
	sortOrder := "ASC"
	if strings.EqualFold(params.SortOrder, "desc") {
		sortOrder = "DESC"
	}

	// Pagination.
	page := params.Page
	if page < 1 {
		page = 1
	}
	pageSize := params.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := fmt.Sprintf(`
		SELECT e.id, e.name, e.email, e.age, e.salary, e.department_id, e.created_at
		%s
		%s
		ORDER BY %s %s
		LIMIT $%d OFFSET $%d`,
		baseFrom, whereClause, sortColumn, sortOrder, argPos, argPos+1)

	args = append(args, pageSize, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("search employees failed: %w", err)
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Salary, &e.DepartmentID, &e.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("search employees failed: %w", err)
		}
		employees = append(employees, e)
	}

	return employees, total, rows.Err()
}

// UpdateEmployeeSalary changes an employee's salary and records the
// change in salary_history - all inside a single database
// transaction, so either both things happen, or neither does.
func UpdateEmployeeSalary(db *sql.DB, employeeID int, newSalary float64) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}
	// If we return before tx.Commit() succeeds, this rolls everything
	// back. Once Commit() has already succeeded, this call is just a
	// harmless no-op.
	defer tx.Rollback()

	var oldSalary float64
	// "FOR UPDATE" locks this specific row until the transaction ends,
	// so two salary updates for the same employee can't race each
	// other and step on one another's changes.
	err = tx.QueryRow(`SELECT salary FROM new_employee WHERE id = $1 FOR UPDATE`, employeeID).Scan(&oldSalary)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("update salary failed: %w", ErrEmployeeNotFound)
		}
		return fmt.Errorf("update salary failed: %w", err)
	}

	if _, err := tx.Exec(`UPDATE new_employee SET salary = $1 WHERE id = $2`, newSalary, employeeID); err != nil {
		return fmt.Errorf("update salary failed: %w", err)
	}

	if _, err := tx.Exec(
		`INSERT INTO salary_history (employee_id, old_salary, new_salary) VALUES ($1, $2, $3)`,
		employeeID, oldSalary, newSalary,
	); err != nil {
		return fmt.Errorf("record salary history failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

// GetSalaryHistory returns every salary change recorded for one
// employee, most recent first.
func GetSalaryHistory(db *sql.DB, employeeID int) ([]model.SalaryHistory, error) {
	query := `
		SELECT id, employee_id, old_salary, new_salary, changed_at
		FROM salary_history
		WHERE employee_id = $1
		ORDER BY changed_at DESC`
	rows, err := db.Query(query, employeeID)
	if err != nil {
		return nil, fmt.Errorf("get salary history failed: %w", err)
	}
	defer rows.Close()

	var history []model.SalaryHistory
	for rows.Next() {
		var h model.SalaryHistory
		if err := rows.Scan(&h.ID, &h.EmployeeID, &h.OldSalary, &h.NewSalary, &h.ChangedAt); err != nil {
			return nil, fmt.Errorf("get salary history failed: %w", err)
		}
		history = append(history, h)
	}
	return history, rows.Err()
}
