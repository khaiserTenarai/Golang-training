package repository

import "customer-management/model"

type CustomerRepository interface {
	AddCustomer(customer model.Customer) error
	GetCustomerByID(id int) (model.Customer, error)
	GetAllCustomers(limit int, offset int) ([]model.Customer, error)
	UpdateCustomer(customer model.Customer) error
	DeleteCustomer(id int) error
	SearchCustomers(keyword string, limit int, offset int) ([]model.Customer, error)
	CountCustomers() (int, error)
}
