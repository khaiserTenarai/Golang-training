package repository
import (
	"context"
	"errors"

	"cms/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(
	db *pgxpool.Pool,
) CustomerRepository {

	return &CustomerRepositoryImpl{
		db: db,
	}
}

func (r *CustomerRepositoryImpl) Save(
	customer model.Customer,
) error {

	query := `
		INSERT INTO customers
		(name, age, email, phone)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Age,
		customer.Email,
		customer.Phone,
	)

	return err
}

func (r *CustomerRepositoryImpl) FindByID(
	id int,
) (model.Customer, error) {

	var customer model.Customer

	query := `
		SELECT id, name, age, email, phone
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
		&customer.Age,
		&customer.Email,
		&customer.Phone,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return customer, errors.New("customer not found")
	}

	return customer, err
}

func (r *CustomerRepositoryImpl) FindAll() (
	[]model.Customer,
	error,
) {

	query := `
		SELECT id, name, age, email, phone
		FROM customers
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {

		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Age,
			&customer.Email,
			&customer.Phone,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, rows.Err()
}

func (r *CustomerRepositoryImpl) Update(
	customer model.Customer,
) error {

	query := `
		UPDATE customers
		SET name = $1,
		    age = $2,
		    email = $3,
		    phone = $4
		WHERE id = $5
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		customer.Name,
		customer.Age,
		customer.Email,
		customer.Phone,
		customer.ID,
	)

	return err
}

func (r *CustomerRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM customers
		WHERE id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}

func (r *CustomerRepositoryImpl) Search(
	keyword string,
) ([]model.Customer, error) {

	query := `
		SELECT id, name, age, email, phone
		FROM customers
		WHERE name ILIKE $1
		   OR email ILIKE $1
		   OR phone ILIKE $1
		ORDER BY id
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		"%"+keyword+"%",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {

		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Age,
			&customer.Email,
			&customer.Phone,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, rows.Err()
}

func (r *CustomerRepositoryImpl) FindWithPagination(
	page int,
	limit int,
) ([]model.Customer, error) {

	offset := (page - 1) * limit

	query := `
		SELECT id, name, age, email, phone
		FROM customers
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(
		context.Background(),
		query,
		limit,
		offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := make([]model.Customer, 0)

	for rows.Next() {

		var customer model.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Age,
			&customer.Email,
			&customer.Phone,
		)

		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, rows.Err()
}
