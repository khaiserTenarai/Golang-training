package service
import "product_inventory/model"

type ProductService interface {

	Save(product model.Product) error

	FindByID(id int) (model.Product, error)

	FindAll() ([]model.Product, error)

	Update(product model.Product) error

	Delete(id int) error

	IncreaseStock(id int, quantity int) error

	DecreaseStock(id int, quantity int) error

	FindLowStock(limit int) ([]model.Product, error)
}
