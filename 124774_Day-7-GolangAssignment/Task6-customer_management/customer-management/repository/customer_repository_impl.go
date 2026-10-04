package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"customer-management/model"
)

var ErrCustomerNotFound = errors.New(
	"customer not found",
)

type CustomerRepositoryImpl struct {
	db *pgxpool.Pool
}

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewCustomerRepository(
	db *pgxpool.Pool,
) CustomerRepository {

	return &CustomerRepositoryImpl{
		db: db,
	}
}

// ==================================================
// CREATE
// ==================================================

func (r *CustomerRepositoryImpl) Save(
	customer model.Customer,
) error {

	query := `
		INSERT INTO customers
			(name, email, phone, city)
		VALUES
			($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.City,
	)

	return err
}

// ==================================================
// FIND BY ID
// ==================================================

func (r *CustomerRepositoryImpl) FindByID(
	id int,
) (model.Customer, error) {

	var customer model.Customer

	query := `
		SELECT
			id,
			name,
			email,
			phone,
			city
		FROM customers
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.City,
	)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return customer, ErrCustomerNotFound
		}

		return customer, err
	}

	return customer, nil
}

// ==================================================
// FIND ALL
// ==================================================

func (r *CustomerRepositoryImpl) FindAll() []model.Customer {

	query := `
		SELECT
			id,
			name,
			email,
			phone,
			city
		FROM customers
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {

		fmt.Println(
			"Error fetching customers:",
			err,
		)

		return []model.Customer{}
	}

	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {

		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
		)

		if err != nil {

			fmt.Println(
				"Error scanning customer:",
				err,
			)

			return []model.Customer{}
		}

		customers = append(
			customers,
			customer,
		)
	}

	return customers
}

// ==================================================
// UPDATE
// ==================================================

func (r *CustomerRepositoryImpl) Update(
	customer model.Customer,
) error {

	query := `
		UPDATE customers
		SET
			name = $1,
			email = $2,
			phone = $3,
			city = $4
		WHERE id = $5
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.City,
		customer.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrCustomerNotFound
	}

	return nil
}

// ==================================================
// DELETE
// ==================================================

func (r *CustomerRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM customers
		WHERE id = $1
	`

	result, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrCustomerNotFound
	}

	return nil
}

// ==================================================
// SEARCH
// ==================================================

func (r *CustomerRepositoryImpl) Search(
	keyword string,
) []model.Customer {

	query := `
		SELECT
			id,
			name,
			email,
			phone,
			city
		FROM customers
		WHERE
			name ILIKE $1
			OR email ILIKE $1
			OR phone ILIKE $1
			OR city ILIKE $1
		ORDER BY id
	`

	searchKeyword := "%" + keyword + "%"

	rows, err := r.db.Query(
		context.Background(),
		query,
		searchKeyword,
	)

	if err != nil {

		fmt.Println(
			"Error searching customers:",
			err,
		)

		return []model.Customer{}
	}

	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {

		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
		)

		if err != nil {

			fmt.Println(
				"Error scanning customer:",
				err,
			)

			return []model.Customer{}
		}

		customers = append(
			customers,
			customer,
		)
	}

	return customers
}

// ==================================================
// PAGINATION
// ==================================================

func (r *CustomerRepositoryImpl) FindPage(
	page int,
	pageSize int,
) []model.Customer {

	offset := (page - 1) * pageSize

	query := `
		SELECT
			id,
			name,
			email,
			phone,
			city
		FROM customers
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		pageSize,
		offset,
	)

	if err != nil {

		fmt.Println(
			"Error fetching page:",
			err,
		)

		return []model.Customer{}
	}

	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {

		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.City,
		)

		if err != nil {

			fmt.Println(
				"Error scanning customer:",
				err,
			)

			return []model.Customer{}
		}

		customers = append(
			customers,
			customer,
		)
	}

	return customers
}
