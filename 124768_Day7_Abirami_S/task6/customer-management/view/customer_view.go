package view

import "customer-management/model"

type CustomerView interface {
	Start()
	CreateCustomer()
	GetCustomer()
	GetAllCustomers()
	UpdateCustomer()
	DeleteCustomer()
	SearchCustomers()
	DisplayCustomer(customer *model.Customer)
	DisplayCustomers(customers []model.Customer)
}
