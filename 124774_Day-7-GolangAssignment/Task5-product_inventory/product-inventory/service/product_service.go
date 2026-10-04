package service

import "prodcut-inventory/model"

type ProductService interface {
	AddProduct(product model.Product) error

	FindProductByID(id int) (model.Product, error)

	FindAllProducts() []model.Product

	UpdateProduct(product model.Product) error

	DeleteProduct(id int) error

	IncreaseStock(id int, quantity int) error

	DecreaseStock(id int, quantity int) error

	FindLowStock() []model.Product
}
