package repository

import "prodcut-inventory/model"

type ProductRepository interface {
	Save(product model.Product) error

	FindByID(id int) (model.Product, error)

	FindAll() []model.Product

	Update(product model.Product) error

	Delete(id int) error

	IncreaseStock(id int, quantity int) error

	DecreaseStock(id int, quantity int) error

	FindLowStock() []model.Product
}
