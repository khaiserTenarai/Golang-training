package repository

import (
	"context"

	"employee-app/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(
	db *pgxpool.Pool,
) EmployeeRepository {

	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) Save(
	employee model.Employee,
) error {

	_, err := r.db.Exec(
		context.Background(),
		`INSERT INTO employees (name, salary)
		 VALUES ($1, $2)`,
		employee.Name,
		employee.Salary,
	)

	return err
}

func (r *EmployeeRepositoryImpl) FindAll() []model.Employee {

	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, name, salary
		 FROM employees
		 ORDER BY id`,
	)

	if err != nil {
		return nil
	}

	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {

		var employee model.Employee

		rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Salary,
		)

		employees = append(
			employees,
			employee,
		)
	}

	return employees
}
