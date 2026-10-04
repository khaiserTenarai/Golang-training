package service

import "customer-management/model"

type CustomerService interface {
	AddCustomer(customer model.Customer) error

	FindCustomerByID(id int) (model.Customer, error)

	FindAllCustomers() []model.Customer

	UpdateCustomer(customer model.Customer) error

	DeleteCustomer(id int) error

	SearchCustomers(keyword string) []model.Customer

	FindCustomersByPage(
		page int,
		pageSize int,
	) []model.Customer
}
