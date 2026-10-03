package repository

import (
	"context"

	"assignment_task12/model"

	"github.com/jackc/pgx/v5"
)

type EmployeeRepositoryImpl struct {
	conn *pgx.Conn
}

func NewEmployeeRepository(
	conn *pgx.Conn,
) EmployeeRepository {

	return &EmployeeRepositoryImpl{

		conn: conn,
	}

}

func (repo *EmployeeRepositoryImpl) Create(
	employee model.Employee,
) error {

	_, err := repo.conn.Exec(

		context.Background(),

		`
		INSERT INTO employees
		(name,email,salary)

		VALUES($1,$2,$3)
		`,

		employee.Name,

		employee.Email,

		employee.Salary,
	)

	return err

}

func (repo *EmployeeRepositoryImpl) GetAll() ([]model.Employee, error) {

	rows, err := repo.conn.Query(

		context.Background(),

		`
		SELECT id,name,email,salary

		FROM employees

		ORDER BY id
		`,
	)

	if err != nil {

		return nil, err

	}

	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {

		var emp model.Employee

		err := rows.Scan(

			&emp.ID,

			&emp.Name,

			&emp.Email,

			&emp.Salary,
		)

		if err != nil {

			return nil, err

		}

		employees = append(employees, emp)

	}

	return employees, nil

}
