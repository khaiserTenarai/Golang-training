package employee

import (
    "context"

    "github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct { DB *pgxpool.Pool }

func NewRepository(databaseURL string) (*Repository, error) {
    pool, err := pgxpool.New(context.Background(), databaseURL)
    if err != nil { return nil, err }
    return &Repository{DB: pool}, nil
}
func (r *Repository) Close() { r.DB.Close() }
func (r *Repository) Ping(ctx context.Context) error { return r.DB.Ping(ctx) }

func (r *Repository) Create(ctx context.Context, e Employee) (Employee, error) {
    err := r.DB.QueryRow(ctx, `
        INSERT INTO employees (name, email, department_id, salary)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, e.Name, e.Email, e.DepartmentID, e.Salary).Scan(&e.ID)
    return e, err
}

func (r *Repository) List(ctx context.Context) ([]Employee, error) {
    rows, err := r.DB.Query(ctx, `
        SELECT id, name, email, department_id, salary
        FROM employees ORDER BY id
    `)
    if err != nil { return nil, err }
    defer rows.Close()
    result := []Employee{}
    for rows.Next() {
        var e Employee
        if err := rows.Scan(&e.ID, &e.Name, &e.Email, &e.DepartmentID, &e.Salary); err != nil {
            return nil, err
        }
        result = append(result, e)
    }
    return result, rows.Err()
}
