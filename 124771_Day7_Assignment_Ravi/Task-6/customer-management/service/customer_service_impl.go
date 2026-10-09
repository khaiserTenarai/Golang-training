package service

import (
	"errors"
	"strings"

	"customer-management/model"
	"customer-management/repository"
)

type CustomerServiceImpl struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{
		customerRepository: customerRepository,
	}
}

func (s *CustomerServiceImpl) AddCustomer(customer model.Customer) error {
	if err := validateCustomer(customer); err != nil {
		return err
	}
	return s.customerRepository.AddCustomer(customer)
}

func (s *CustomerServiceImpl) GetCustomerByID(id int) (model.Customer, error) {
	if id <= 0 {
		return model.Customer{}, errors.New("ID must be greater than 0")
	}
	return s.customerRepository.GetCustomerByID(id)
}

func (s *CustomerServiceImpl) GetAllCustomers(limit int, offset int) ([]model.Customer, error) {
	if limit <= 0 {
		limit = 5
	}
	if offset < 0 {
		offset = 0
	}
	return s.customerRepository.GetAllCustomers(limit, offset)
}

func (s *CustomerServiceImpl) UpdateCustomer(customer model.Customer) error {
	if customer.ID <= 0 {
		return errors.New("ID must be greater than 0")
	}
	if err := validateCustomer(customer); err != nil {
		return err
	}
	return s.customerRepository.UpdateCustomer(customer)
}

func (s *CustomerServiceImpl) DeleteCustomer(id int) error {
	if id <= 0 {
		return errors.New("ID must be greater than 0")
	}
	return s.customerRepository.DeleteCustomer(id)
}

func (s *CustomerServiceImpl) SearchCustomers(keyword string, limit int, offset int) ([]model.Customer, error) {
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		return nil, errors.New("search keyword cannot be empty")
	}
	if limit <= 0 {
		limit = 5
	}
	if offset < 0 {
		offset = 0
	}
	return s.customerRepository.SearchCustomers(keyword, limit, offset)
}

func validateCustomer(customer model.Customer) error {
	if strings.TrimSpace(customer.Name) == "" {
		return errors.New("customer name cannot be empty")
	}
	if strings.TrimSpace(customer.Email) == "" {
		return errors.New("customer email cannot be empty")
	}
	if strings.TrimSpace(customer.Phone) == "" {
		return errors.New("customer phone cannot be empty")
	}
	return nil
}
