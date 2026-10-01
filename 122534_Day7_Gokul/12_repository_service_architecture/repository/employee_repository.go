package repository

import (
	"context"
	"example.com/q12-repository-service/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository interface {
	Create(context.Context, model.Employee) error
	FindAll(context.Context) ([]model.Employee, error)
}
type PostgresEmployeeRepository struct{ db *pgxpool.Pool }

func NewPostgresEmployeeRepository(db *pgxpool.Pool) *PostgresEmployeeRepository {
	return &PostgresEmployeeRepository{db: db}
}
func (r *PostgresEmployeeRepository) Create(c context.Context, e model.Employee) error {
	_, err := r.db.Exec(c, `INSERT INTO employees(name,email,department) VALUES($1,$2,$3)`, e.Name, e.Email, e.Department)
	return err
}
func (r *PostgresEmployeeRepository) FindAll(c context.Context) ([]model.Employee, error) {
	rows, err := r.db.Query(c, `SELECT id,name,email,department FROM employees ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var list []model.Employee
	for rows.Next() {
		var e model.Employee
		if err = rows.Scan(&e.ID, &e.Name, &e.Email, &e.Department); err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, rows.Err()
}
