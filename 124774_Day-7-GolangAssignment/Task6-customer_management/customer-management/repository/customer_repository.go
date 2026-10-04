package repository

import "customer-management/model"

type CustomerRepository interface {
	Save(customer model.Customer) error

	FindByID(id int) (model.Customer, error)

	FindAll() []model.Customer

	Update(customer model.Customer) error

	Delete(id int) error

	Search(keyword string) []model.Customer

	FindPage(page int, pageSize int) []model.Customer
}
