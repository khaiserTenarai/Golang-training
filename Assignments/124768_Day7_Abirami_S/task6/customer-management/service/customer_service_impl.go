package service

import (
	"customer-management/model"
	"customer-management/repository"
)

type CustomerServiceImpl struct {
	repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		repository: repository,
	}
}

func (s *CustomerServiceImpl) CreateCustomer(customer model.Customer) error {
	return s.repository.CreateCustomer(customer)
}

func (s *CustomerServiceImpl) GetCustomer(id int) (*model.Customer, error) {
	return s.repository.GetCustomer(id)
}

func (s *CustomerServiceImpl) GetAllCustomers() ([]model.Customer, error) {
	return s.repository.GetAllCustomers()
}

func (s *CustomerServiceImpl) UpdateCustomer(customer model.Customer) error {
	return s.repository.UpdateCustomer(customer)
}

func (s *CustomerServiceImpl) DeleteCustomer(id int) error {
	return s.repository.DeleteCustomer(id)
}

func (s *CustomerServiceImpl) SearchCustomers(name string, email string, city string, page int, size int) ([]model.Customer, error) {
	return s.repository.SearchCustomers(name, email, city, page, size)
}
