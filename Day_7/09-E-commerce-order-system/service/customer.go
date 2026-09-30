package service
import (
	"ecommerce/model"
	"ecommerce/repository"
	"ecommerce/utility"
)

type CustomerService interface {
	Save(customer model.Customer) error
	FindByID(id int) (model.Customer, error)
	FindAll() ([]model.Customer, error)
	Update(customer model.Customer) error
	Delete(id int) error
}

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

	return s.repository.Save(customer)
}

func (s *CustomerServiceImpl) FindByID(
	id int,
) (model.Customer, error) {

	err := utility.ValidateID(id)

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

	err := utility.ValidateID(customer.ID)

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

	return s.repository.Update(customer)
}

func (s *CustomerServiceImpl) Delete(id int) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}
