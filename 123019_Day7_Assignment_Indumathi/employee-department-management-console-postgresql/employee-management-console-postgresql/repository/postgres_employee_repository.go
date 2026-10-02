package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"example.com/employee-management/models"
)

var ErrNotFound = errors.New("employee not found")

type PostgresEmployeeRepository struct {
	db *pgxpool.Pool
}

// Constructor
func NewPostgresEmployeeRepository(
	db *pgxpool.Pool,
) *PostgresEmployeeRepository {
	return &PostgresEmployeeRepository{
		db: db,
	}
}

// ==================================================
// CREATE
// ==================================================

func (r *PostgresEmployeeRepository) Save(
	employee models.Employee,
) error {
	fmt.Println("\n----- CREATE EMPLOYEE -----")

	query := `
		INSERT INTO employees
		(name, email, age, salary, city, state, pincode, department_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.City,
		employee.State,
		employee.Pincode,
		employee.DepartmentID, // Can be nil or &int
	).Scan(&employee.ID)

	if err != nil {
		return err
	}

	fmt.Println("Employee created successfully")
	fmt.Println("Generated ID:", employee.ID)

	return nil
}

// ==================================================
// READ ONE
// ==================================================

func (r *PostgresEmployeeRepository) FindByID(
	id int64,
) (models.Employee, error) {
	fmt.Println("\n----- READ EMPLOYEE -----")

	var employee models.Employee

	query := `
		SELECT
			id,
			name,
			email,
			age,
			salary,
			city,
			state,
			pincode,
			department_id
		FROM employees
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Age,
		&employee.Salary,
		&employee.City,
		&employee.State,
		&employee.Pincode,
		&employee.DepartmentID,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return employee, fmt.Errorf("employee not found")
		}
		return employee, err
	}

	return employee, nil
}

// ==================================================
// READ ALL
// ==================================================

func (r *PostgresEmployeeRepository) FindAll() (
	[]models.Employee,
	error,
) {
	fmt.Println("\n----- READ ALL EMPLOYEES -----")

	query := `
		SELECT
			id,
			name,
			email,
			age,
			salary,
			city,
			state,
			pincode,
			department_id
		FROM employees
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Age,
			&employee.Salary,
			&employee.City,
			&employee.State,
			&employee.Pincode,
			&employee.DepartmentID,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

// ==================================================
// UPDATE
// ==================================================

func (r *PostgresEmployeeRepository) Update(
	employee models.Employee,
) error {
	fmt.Println("\n----- UPDATE EMPLOYEE -----")

	query := `
		UPDATE employees
		SET
			name = $1,
			email = $2,
			age = $3,
			salary = $4,
			city = $5,
			state = $6,
			pincode = $7,
			department_id = $8
		WHERE id = $9
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		employee.Name,
		employee.Email,
		employee.Age,
		employee.Salary,
		employee.City,
		employee.State,
		employee.Pincode,
		employee.DepartmentID,
		employee.ID,
	)

	if err != nil {
		return err
	}

	fmt.Println("Employee updated successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

// ==================================================
// DELETE
// ==================================================

func (r *PostgresEmployeeRepository) Delete(
	id int64,
) error {
	fmt.Println("\n----- DELETE EMPLOYEE -----")

	result, err := r.db.Exec(
		context.Background(),
		`DELETE FROM employees WHERE id = $1`,
		id,
	)

	if err != nil {
		return err
	}

	fmt.Println("Employee deleted successfully")
	fmt.Println("Rows affected:", result.RowsAffected())

	return nil
}

// ==================================================
// TASK 2: DEPARTMENT MAPPING
// ==================================================

// AssignDepartment updates the foreign key department_id for an employee
func (r *PostgresEmployeeRepository) AssignDepartment(
	employeeID int64,
	departmentID int,
) error {
	fmt.Println("\n----- ASSIGN DEPARTMENT TO EMPLOYEE -----")

	query := `
		UPDATE employees
		SET department_id = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	res, err := r.db.Exec(context.Background(), query, departmentID, employeeID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("employee not found")
	}

	fmt.Println("Department assigned successfully")
	return nil
}

// FindWithDepartment fetches employee along with department name and code using JOIN
// FindWithDepartment fetches employee along with department name and code using JOIN
func (r *PostgresEmployeeRepository) FindWithDepartment(
	id int64,
) (models.EmployeeWithDepartment, error) {
	fmt.Println("\n----- READ EMPLOYEE WITH DEPARTMENT -----")

	query := `
		SELECT 
			e.id, e.name, e.email, e.age, e.salary, e.city, e.state, e.pincode, e.department_id,
			COALESCE(d.name, 'Unassigned') as department_name, 
			COALESCE(d.code, 'N/A') as department_code
		FROM employees e
		LEFT JOIN departments d ON e.department_id = d.id
		WHERE e.id = $1
	`

	var emp models.EmployeeWithDepartment
	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&emp.ID, &emp.Name, &emp.Email, &emp.Age, &emp.Salary, &emp.City, &emp.State, &emp.Pincode, &emp.DepartmentID,
		&emp.DepartmentName, &emp.DepartmentCode,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return emp, fmt.Errorf("employee not found")
		}
		return emp, err
	}

	return emp, nil
}