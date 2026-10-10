// Package repository contains database access logic.
package repository

// Import context for request-scoped database operations.
import "context"

// Import errors for not-found errors.
import "errors"

// Import pgxpool for PostgreSQL connection pooling.
import "github.com/jackc/pgx/v5/pgxpool"

// Import the employee model.
import "employee-management/model"

// EmployeeRepository defines employee database operations.
type EmployeeRepository interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context, string, string, string, string, int, int) ([]model.Employee, int, error)
	GetByID(context.Context, int64) (*model.Employee, error)
	Update(context.Context, int64, *model.Employee) (*model.Employee, error)
	Delete(context.Context, int64) error
	Ping(context.Context) error
}

// PostgresEmployeeRepository implements EmployeeRepository with PostgreSQL.
type PostgresEmployeeRepository struct{ db *pgxpool.Pool }

// NewPostgresEmployeeRepository creates a PostgreSQL repository.
func NewPostgresEmployeeRepository(db *pgxpool.Pool) EmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}

// Create inserts an employee.
func (r *PostgresEmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	return r.db.QueryRow(ctx, `INSERT INTO employees (name,email,department,salary) VALUES ($1,$2,$3,$4) RETURNING id`, e.Name, e.Email, e.Department, e.Salary).Scan(&e.ID)
}

// GetAll returns filtered, sorted and paginated employees.
func (r *PostgresEmployeeRepository) GetAll(ctx context.Context, department, name, sortBy, sortOrder string, page, pageSize int) ([]model.Employee, int, error) {
	orderBy := "id"
	switch sortBy {
	case "name":
		orderBy = "name"
	case "email":
		orderBy = "email"
	case "department":
		orderBy = "department"
	case "salary":
		orderBy = "salary"
	}
	direction := "ASC"
	if sortOrder == "desc" {
		direction = "DESC"
	}
	countSQL := `SELECT COUNT(*) FROM employees WHERE ($1='' OR department=$1) AND ($2='' OR name ILIKE '%'||$2||'%')`
	var total int
	if err := r.db.QueryRow(ctx, countSQL, department, name).Scan(&total); err != nil {
		return nil, 0, err
	}
	offset := (page - 1) * pageSize
	query := `SELECT id,name,email,department,salary FROM employees WHERE ($1='' OR department=$1) AND ($2='' OR name ILIKE '%'||$2||'%') ORDER BY ` + orderBy + ` ` + direction + ` LIMIT $3 OFFSET $4`
	rows, err := r.db.Query(ctx, query, department, name, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	employees := make([]model.Employee, 0)
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary); err != nil {
			return nil, 0, err
		}
		employees = append(employees, e)
	}
	return employees, total, rows.Err()
}

// GetByID returns one employee.
func (r *PostgresEmployeeRepository) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	var e model.Employee
	err := r.db.QueryRow(ctx, `SELECT id,name,email,department,salary FROM employees WHERE id=$1`, id).Scan(&e.ID, &e.Name, &e.Email, &e.Department, &e.Salary)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// Update modifies an employee.
func (r *PostgresEmployeeRepository) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	var u model.Employee
	err := r.db.QueryRow(ctx, `UPDATE employees SET name=$1,email=$2,department=$3,salary=$4 WHERE id=$5 RETURNING id,name,email,department,salary`, e.Name, e.Email, e.Department, e.Salary, id).Scan(&u.ID, &u.Name, &u.Email, &u.Department, &u.Salary)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Delete removes an employee.
func (r *PostgresEmployeeRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.Exec(ctx, `DELETE FROM employees WHERE id=$1`, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return errors.New("employee not found")
	}
	return nil
}

// Ping checks PostgreSQL availability.
func (r *PostgresEmployeeRepository) Ping(ctx context.Context) error { return r.db.Ping(ctx) }
