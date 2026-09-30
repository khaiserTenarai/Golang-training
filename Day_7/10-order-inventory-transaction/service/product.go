package service
import (
	"order_inventory/model"
	"order_inventory/repository"
	"order_inventory/utility"
)

type ProductService interface {
	Save(product model.Product) error
	FindByID(id int) (model.Product, error)
	FindAll() ([]model.Product, error)
	Update(product model.Product) error
}

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

func NewProductService(
	repository repository.ProductRepository,
) ProductService {

	return &ProductServiceImpl{
		repository: repository,
	}
}

func (s *ProductServiceImpl) Save(
	product model.Product,
) error {

	err := utility.ValidateProduct(product)

	if err != nil {
		return err
	}

	return s.repository.Save(product)
}

func (s *ProductServiceImpl) FindByID(
	id int,
) (model.Product, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return model.Product{}, err
	}

	return s.repository.FindByID(id)
}

func (s *ProductServiceImpl) FindAll() (
	[]model.Product,
	error,
) {

	return s.repository.FindAll()
}

func (s *ProductServiceImpl) Update(
	product model.Product,
) error {

	err := utility.ValidateID(product.ID)

	if err != nil {
		return err
	}

	err = utility.ValidateProduct(product)

	if err != nil {
		return err
	}

	return s.repository.Update(product)
}
