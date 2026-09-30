package repository
import (
	"context"
	"errors"

	"ecommerce/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository interface {
	Save(product model.Product) error
	FindByID(id int) (model.Product, error)
	FindAll() ([]model.Product, error)
	Update(product model.Product) error
	Delete(id int) error
}

type ProductRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewProductRepository(
	db *pgxpool.Pool,
) ProductRepository {

	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) Save(
	product model.Product,
) error {

	query := `
		INSERT INTO products
		(name, price, stock)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
	)

	return err
}

func (r *ProductRepositoryImpl) FindByID(
	id int,
) (model.Product, error) {

	var product model.Product

	query := `
		SELECT id, name, price, stock
		FROM products
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Stock,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return product, errors.New("product not found")
	}

	return product, err
}

func (r *ProductRepositoryImpl) FindAll() (
	[]model.Product,
	error,
) {

	query := `
		SELECT id, name, price, stock
		FROM products
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

	products := make([]model.Product, 0)

	for rows.Next() {

		var product model.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, rows.Err()
}

func (r *ProductRepositoryImpl) Update(
	product model.Product,
) error {

	query := `
		UPDATE products
		SET name = $1,
		    price = $2,
		    stock = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		product.Name,
		product.Price,
		product.Stock,
		product.ID,
	)

	return err
}

func (r *ProductRepositoryImpl) Delete(
	id int,
) error {

	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}
