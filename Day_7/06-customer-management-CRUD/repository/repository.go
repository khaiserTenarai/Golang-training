package repository
import "cms/model"

type CustomerRepository interface {

	Save(customer model.Customer) error

	FindByID(id int) (model.Customer, error)

	FindAll() ([]model.Customer, error)

	Update(customer model.Customer) error

	Delete(id int) error

	Search(keyword string) ([]model.Customer, error)

	FindWithPagination(page int, limit int) ([]model.Customer, error)
}
