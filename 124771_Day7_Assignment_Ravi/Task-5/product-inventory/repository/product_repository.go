package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"product-inventory/model"
)

type ProductRepository interface {
	Create(
		ctx context.Context,
		product *model.Product,
	) (*model.Product, error)

	GetByID(
		ctx context.Context,
		id int64,
	) (*model.Product, error)

	GetAll(
		ctx context.Context,
	) ([]model.Product, error)

	Update(
		ctx context.Context,
		product *model.Product,
	) (*model.Product, error)

	Delete(
		ctx context.Context,
		id int64,
	) error

	GetByIDForUpdate(
		ctx context.Context,
		tx pgx.Tx,
		id int64,
	) (*model.Product, error)

	UpdateStock(
		ctx context.Context,
		tx pgx.Tx,
		id int64,
		stockQuantity int,
	) error

	GetLowStockProducts(
		ctx context.Context,
	) ([]model.Product, error)
}
