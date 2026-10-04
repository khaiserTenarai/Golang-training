package service

import (
	"strings"

	"customer-management/model"
	"customer-management/repository"
	"customer-management/utility"
)

type CustomerServiceImpl struct {
	repository repository.CustomerRepository
}

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewCustomerService(
	repository repository.CustomerRepository,
) CustomerService {

	return &CustomerServiceImpl{
		repository: repository,
	}
}

// ==================================================
// CREATE
// ==================================================

func (s *CustomerServiceImpl) AddCustomer(
	customer model.Customer,
) error {

	if err := utility.ValidateName(
		customer.Name,
	); err != nil {
		return err
	}

	if err := utility.ValidateEmail(
		customer.Email,
	); err != nil {
		return err
	}

	if err := utility.ValidatePhone(
		customer.Phone,
	); err != nil {
		return err
	}

	if err := utility.ValidateCity(
		customer.City,
	); err != nil {
		return err
	}

	return s.repository.Save(customer)
}

// ==================================================
// FIND BY ID
// ==================================================

func (s *CustomerServiceImpl) FindCustomerByID(
	id int,
) (model.Customer, error) {

	if err := utility.ValidateID(id); err != nil {
		return model.Customer{}, err
	}

	return s.repository.FindByID(id)
}

// ==================================================
// FIND ALL
// ==================================================

func (s *CustomerServiceImpl) FindAllCustomers() []model.Customer {

	return s.repository.FindAll()
}

// ==================================================
// UPDATE
// ==================================================

func (s *CustomerServiceImpl) UpdateCustomer(
	customer model.Customer,
) error {

	if err := utility.ValidateID(
		customer.ID,
	); err != nil {
		return err
	}

	if err := utility.ValidateName(
		customer.Name,
	); err != nil {
		return err
	}

	if err := utility.ValidateEmail(
		customer.Email,
	); err != nil {
		return err
	}

	if err := utility.ValidatePhone(
		customer.Phone,
	); err != nil {
		return err
	}

	if err := utility.ValidateCity(
		customer.City,
	); err != nil {
		return err
	}

	return s.repository.Update(customer)
}

// ==================================================
// DELETE
// ==================================================

func (s *CustomerServiceImpl) DeleteCustomer(
	id int,
) error {

	if err := utility.ValidateID(id); err != nil {
		return err
	}

	return s.repository.Delete(id)
}

// ==================================================
// SEARCH
// ==================================================

func (s *CustomerServiceImpl) SearchCustomers(
	keyword string,
) []model.Customer {

	if strings.TrimSpace(keyword) == "" {
		return []model.Customer{}
	}

	return s.repository.Search(keyword)
}

// ==================================================
// PAGINATION
// ==================================================

func (s *CustomerServiceImpl) FindCustomersByPage(
	page int,
	pageSize int,
) []model.Customer {

	if err := utility.ValidatePage(page); err != nil {
		return []model.Customer{}
	}

	if err := utility.ValidatePageSize(pageSize); err != nil {
		return []model.Customer{}
	}

	return s.repository.FindPage(
		page,
		pageSize,
	)
}
