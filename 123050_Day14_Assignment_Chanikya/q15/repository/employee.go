package repository

import (
 "context"
 "database/sql"
 "errors"
 "fmt"

 "employee-management/model"
)

var ErrNotFound = errors.New("employee not found")

type EmployeeRepository interface {
 Create(context.Context, model.Employee) (model.Employee, error)
 List(context.Context, int, int, string, string) ([]model.Employee, int64, error)
 Get(context.Context, int64) (model.Employee, error)
 Update(context.Context, int64, model.Employee) (model.Employee, error)
 Delete(context.Context, int64) error
}

type postgresRepo struct { db *sql.DB }
func New(db *sql.DB) EmployeeRepository { return &postgresRepo{db: db} }

func scanEmployee(row interface{ Scan(...any) error }) (model.Employee, error) {
 var e model.Employee
 err := row.Scan(&e.ID, &e.Name, &e.Email, &e.Age, &e.Salary, &e.CreatedAt, &e.UpdatedAt)
 if errors.Is(err, sql.ErrNoRows) { return e, ErrNotFound }
 return e, err
}
const returning = "id, name, email, age, salary, created_at, updated_at"

func (r *postgresRepo) Create(ctx context.Context, e model.Employee) (model.Employee, error) {
 q := "INSERT INTO employees(name,email,age,salary) VALUES($1,$2,$3,$4) RETURNING " + returning
 return scanEmployee(r.db.QueryRowContext(ctx, q, e.Name, e.Email, e.Age, e.Salary))
}
func (r *postgresRepo) List(ctx context.Context, page, limit int, search, sort string) ([]model.Employee, int64, error) {
 allowed := map[string]string{"id":"id", "name":"name", "email":"email", "age":"age", "salary":"salary"}
 col, ok := allowed[sort]; if !ok { col = "id" }
 pattern := "%" + search + "%"
 var total int64
 if err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM employees WHERE name ILIKE $1 OR email ILIKE $1", pattern).Scan(&total); err != nil { return nil, 0, err }
 q := fmt.Sprintf("SELECT %s FROM employees WHERE name ILIKE $1 OR email ILIKE $1 ORDER BY %s LIMIT $2 OFFSET $3", returning, col)
 rows, err := r.db.QueryContext(ctx, q, pattern, limit, (page-1)*limit); if err != nil { return nil, 0, err }; defer rows.Close()
 out := make([]model.Employee, 0)
 for rows.Next() { e, err := scanEmployee(rows); if err != nil { return nil, 0, err }; out = append(out, e) }
 if err := rows.Err(); err != nil { return nil, 0, err }
 return out, total, nil
}
func (r *postgresRepo) Get(ctx context.Context, id int64) (model.Employee, error) {
 return scanEmployee(r.db.QueryRowContext(ctx, "SELECT "+returning+" FROM employees WHERE id=$1", id))
}
func (r *postgresRepo) Update(ctx context.Context, id int64, e model.Employee) (model.Employee, error) {
 q := "UPDATE employees SET name=$1,email=$2,age=$3,salary=$4,updated_at=NOW() WHERE id=$5 RETURNING " + returning
 return scanEmployee(r.db.QueryRowContext(ctx, q, e.Name, e.Email, e.Age, e.Salary, id))
}
func (r *postgresRepo) Delete(ctx context.Context, id int64) error {
 res, err := r.db.ExecContext(ctx, "DELETE FROM employees WHERE id=$1", id); if err != nil { return err }
 n, err := res.RowsAffected(); if err != nil { return err }; if n == 0 { return ErrNotFound }; return nil
}
