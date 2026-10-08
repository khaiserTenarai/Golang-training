// Package repository contains database access logic.
package repository

// Import context for request-scoped database operations.
import "context"

// Import pgxpool for PostgreSQL connection pooling.
import "github.com/jackc/pgx/v5/pgxpool"

// Import errors for the delete not-found error.
import "errors"

// Import the employee model.
import "employee-management/model"

// EmployeeRepository defines database operations for employees.
type EmployeeRepository interface {
	// Create inserts a new employee.
	Create(ctx context.Context, employee *model.Employee) error
	// GetAll retrieves all employees.
	GetAll(ctx context.Context, sortBy string) ([]model.Employee, error)
	// GetByID retrieves one employee by path parameter ID.
	GetByID(ctx context.Context, id int64) (*model.Employee, error)
	// Update modifies an existing employee.
	Update(ctx context.Context, id int64, employee *model.Employee) (*model.Employee, error)
	// Delete removes an employee.
	Delete(ctx context.Context, id int64) error
}

// PostgresEmployeeRepository implements EmployeeRepository using PostgreSQL.
type PostgresEmployeeRepository struct {
	// db stores the PostgreSQL connection pool.
	db *pgxpool.Pool
}

// NewPostgresEmployeeRepository creates a PostgreSQL repository.
func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	// Return the repository implementation through the interface.
	return &PostgresEmployeeRepository{db: db}
}

// Create inserts an employee and returns the generated ID.
func (r *PostgresEmployeeRepository) Create(ctx context.Context, employee *model.Employee) error {
	// Execute the INSERT statement and scan the generated ID.
	return r.db.QueryRow(ctx, `INSERT INTO employees (name, email, department, salary) VALUES ($1, $2, $3, $4) RETURNING id`, employee.Name, employee.Email, employee.Department, employee.Salary).Scan(&employee.ID)
}

// GetAll retrieves employees and optionally sorts them by a safe predefined column.
func (r *PostgresEmployeeRepository) GetAll(ctx context.Context, sortBy string) ([]model.Employee, error) {
	// Default to sorting by employee ID.
	orderBy := "id"
	// Allow only predefined sorting columns to avoid SQL injection.
	if sortBy == "name" {
		// Sort by employee name when requested.
		orderBy = "name"
	}
	// Query all employees without filtering or pagination.
	rows, err := r.db.Query(ctx, "SELECT id, name, email, department, salary FROM employees ORDER BY "+orderBy)
	// Return the database error when the query fails.
	if err != nil {
		// Return nil results with the database error.
		return nil, err
	}
	// Ensure database rows are closed after processing.
	defer rows.Close()
	// Create an empty employee slice.
	employees := make([]model.Employee, 0)
	// Iterate through all database rows.
	for rows.Next() {
		// Create an employee variable for the current row.
		var employee model.Employee
		// Scan database columns into the employee structure.
		if err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Department, &employee.Salary); err != nil {
			// Return the scanning error.
			return nil, err
		}
		// Add the employee to the result slice.
		employees = append(employees, employee)
	}
	// Return the complete employee list.
	return employees, rows.Err()
}

// GetByID retrieves an employee by ID.
func (r *PostgresEmployeeRepository) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	// Create an employee variable for the database result.
	var employee model.Employee
	// Query PostgreSQL using the path parameter ID.
	err := r.db.QueryRow(ctx, "SELECT id, name, email, department, salary FROM employees WHERE id = $1", id).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Department, &employee.Salary)
	// Return the employee and any database error.
	if err != nil {
		// Return nil because no valid employee was retrieved.
		return nil, err
	}
	// Return the retrieved employee.
	return &employee, nil
}

// Update updates an employee by ID.
func (r *PostgresEmployeeRepository) Update(ctx context.Context, id int64, employee *model.Employee) (*model.Employee, error) {
	// Create a variable for the updated employee.
	var updated model.Employee
	// Execute the UPDATE statement and return the updated record.
	err := r.db.QueryRow(ctx, "UPDATE employees SET name=$1, email=$2, department=$3, salary=$4 WHERE id=$5 RETURNING id, name, email, department, salary", employee.Name, employee.Email, employee.Department, employee.Salary, id).Scan(&updated.ID, &updated.Name, &updated.Email, &updated.Department, &updated.Salary)
	// Return nil and the error when the employee does not exist or the query fails.
	if err != nil {
		// Return nil because the update did not succeed.
		return nil, err
	}
	// Return the updated employee.
	return &updated, nil
}

// Delete deletes an employee by ID.
func (r *PostgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	// Execute the DELETE statement using the path parameter ID.
	result, err := r.db.Exec(ctx, "DELETE FROM employees WHERE id = $1", id)
	// Return any database execution error.
	if err != nil {
		// Return the database error.
		return err
	}
	// Check whether any employee was deleted.
	if result.RowsAffected() == 0 {
		// Return a descriptive not-found error.
		return errors.New("employee not found")
	}
	// Return nil when deletion succeeds.
	return nil
}
