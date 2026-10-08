package controller

import (
	"customer-management/model"
	"customer-management/service"
)

type CustomerControllerImpl struct {
	service service.CustomerService
}

func NewCustomerController(service service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		service: service,
	}
}

func (c *CustomerControllerImpl) CreateCustomer(customer model.Customer) error {
	return c.service.CreateCustomer(customer)
}

func (c *CustomerControllerImpl) GetCustomer(id int) (*model.Customer, error) {
	return c.service.GetCustomer(id)
}

func (c *CustomerControllerImpl) GetAllCustomers() ([]model.Customer, error) {
	return c.service.GetAllCustomers()
}

func (c *CustomerControllerImpl) UpdateCustomer(customer model.Customer) error {
	return c.service.UpdateCustomer(customer)
}

func (c *CustomerControllerImpl) DeleteCustomer(id int) error {
	return c.service.DeleteCustomer(id)
}

func (c *CustomerControllerImpl) SearchCustomers(name string, email string, city string, page int, size int) ([]model.Customer, error) {
	return c.service.SearchCustomers(name, email, city, page, size)
}
