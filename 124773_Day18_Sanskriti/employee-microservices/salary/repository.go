package salary

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct { DB *pgxpool.Pool }
func NewRepository(url string) (*Repository, error) {
    pool, err := pgxpool.New(context.Background(), url)
    if err != nil { return nil, err }
    return &Repository{DB:pool}, nil
}
func (r *Repository) Close() { r.DB.Close() }
func (r *Repository) Ping(ctx context.Context) error { return r.DB.Ping(ctx) }
func (r *Repository) ListByEmployee(ctx context.Context, employeeID int64) ([]Salary, error) {
    rows, err := r.DB.Query(ctx, "SELECT id, employee_id, amount FROM salaries WHERE employee_id=$1 ORDER BY id", employeeID)
    if err != nil { return nil, err }
    defer rows.Close()
    result := []Salary{}
    for rows.Next() {
        var item Salary
        if err := rows.Scan(&item.ID, &item.EmployeeID, &item.Amount); err != nil { return nil, err }
        result = append(result, item)
    }
    return result, rows.Err()
}
func (r *Repository) Create(ctx context.Context, employeeID int64, amount float64) error {
    _, err := r.DB.Exec(ctx, "INSERT INTO salaries (employee_id, amount) VALUES ($1, $2)", employeeID, amount)
    return err
}
