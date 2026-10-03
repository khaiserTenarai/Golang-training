package repository

import "customer-management/model"

type CustomerRepository interface {
	CreateCustomer(customer model.Customer) error
	GetCustomer(id int) (*model.Customer, error)
	GetAllCustomers() ([]model.Customer, error)
	UpdateCustomer(customer model.Customer) error
	DeleteCustomer(id int) error
	SearchCustomers(name string, email string, city string, page int, size int) ([]model.Customer, error)
}
