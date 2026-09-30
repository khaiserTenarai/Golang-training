package service
import (
	"cms/model"
	"cms/repository"
	"cms/utility"
)

type CustomerServiceImpl struct {
	repository repository.CustomerRepository
}

func NewCustomerService(
	repository repository.CustomerRepository,
) CustomerService {

	return &CustomerServiceImpl{
		repository: repository,
	}
}

func (s *CustomerServiceImpl) Save(
	customer model.Customer,
) error {

	err := utility.ValidateCustomer(customer)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(customer.Email)

	if err != nil {
		return err
	}

	err = utility.ValidatePhone(customer.Phone)

	if err != nil {
		return err
	}

	return s.repository.Save(customer)
}

func (s *CustomerServiceImpl) FindByID(
	id int,
) (model.Customer, error) {

	err := utility.ValidateCustomerID(id)

	if err != nil {
		return model.Customer{}, err
	}

	return s.repository.FindByID(id)
}

func (s *CustomerServiceImpl) FindAll() (
	[]model.Customer,
	error,
) {

	return s.repository.FindAll()
}

func (s *CustomerServiceImpl) Update(
	customer model.Customer,
) error {

	err := utility.ValidateCustomerID(customer.ID)

	if err != nil {
		return err
	}

	err = utility.ValidateCustomer(customer)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(customer.Email)

	if err != nil {
		return err
	}

	err = utility.ValidatePhone(customer.Phone)

	if err != nil {
		return err
	}

	return s.repository.Update(customer)
}

func (s *CustomerServiceImpl) Delete(
	id int,
) error {

	err := utility.ValidateCustomerID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}

func (s *CustomerServiceImpl) Search(
	keyword string,
) ([]model.Customer, error) {

	err := utility.ValidateSearch(keyword)

	if err != nil {
		return nil, err
	}

	return s.repository.Search(keyword)
}

func (s *CustomerServiceImpl) FindWithPagination(
	page int,
	limit int,
) ([]model.Customer, error) {

	err := utility.ValidatePagination(
		page,
		limit,
	)

	if err != nil {
		return nil, err
	}

	return s.repository.FindWithPagination(
		page,
		limit,
	)
}
