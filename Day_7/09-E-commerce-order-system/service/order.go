package service
import (
	"ecommerce/model"
	"ecommerce/repository"
	"ecommerce/utility"
)

type OrderService interface {
	CreateOrder(order model.OrderRequest) error
	FindByID(id int) ([]model.OrderDetail, error)
	FindAll() ([]model.OrderDetail, error)
}

type OrderServiceImpl struct {
	repository repository.OrderRepository
}

func NewOrderService(
	repository repository.OrderRepository,
) OrderService {

	return &OrderServiceImpl{
		repository: repository,
	}
}

func (s *OrderServiceImpl) CreateOrder(
	order model.OrderRequest,
) error {

	err := utility.ValidateOrder(order)

	if err != nil {
		return err
	}

	return s.repository.CreateOrder(order)
}

func (s *OrderServiceImpl) FindByID(
	id int,
) ([]model.OrderDetail, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return nil, err
	}

	return s.repository.FindByID(id)
}

func (s *OrderServiceImpl) FindAll() (
	[]model.OrderDetail,
	error,
) {

	return s.repository.FindAll()
}
