package department

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
func (r *Repository) Get(ctx context.Context, id int64) (Department, error) {
    var d Department
    err := r.DB.QueryRow(ctx, "SELECT id, name FROM departments WHERE id=$1", id).Scan(&d.ID, &d.Name)
    return d, err
}
func (r *Repository) List(ctx context.Context) ([]Department, error) {
    rows, err := r.DB.Query(ctx, "SELECT id, name FROM departments ORDER BY id")
    if err != nil { return nil, err }
    defer rows.Close()
    result := []Department{}
    for rows.Next() {
        var d Department
        if err := rows.Scan(&d.ID, &d.Name); err != nil { return nil, err }
        result = append(result, d)
    }
    return result, rows.Err()
}
