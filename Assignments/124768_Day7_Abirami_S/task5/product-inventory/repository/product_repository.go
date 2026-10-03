package repository

import "product-inventory/model"

type ProductRepository interface {
	CreateProduct(product model.Product) error
	GetProduct(id int) (*model.Product, error)
	GetAllProducts() ([]model.Product, error)
	UpdateProduct(product model.Product) error
	DeleteProduct(id int) error
	IncreaseStock(id int, quantity int) error
	DecreaseStock(id int, quantity int) error
	GetLowStockProducts() ([]model.Product, error)
}
