package service
import (
	"product_inventory/model"
	"product_inventory/repository"
	"product_inventory/utility"
)

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

func (s *ProductServiceImpl) Delete(
	id int,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}

func (s *ProductServiceImpl) IncreaseStock(
	id int,
	quantity int,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	err = utility.ValidateQuantity(quantity)

	if err != nil {
		return err
	}

	return s.repository.IncreaseStock(
		id,
		quantity,
	)
}

func (s *ProductServiceImpl) DecreaseStock(
	id int,
	quantity int,
) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	err = utility.ValidateQuantity(quantity)

	if err != nil {
		return err
	}

	return s.repository.DecreaseStock(
		id,
		quantity,
	)
}

func (s *ProductServiceImpl) FindLowStock(
	limit int,
) ([]model.Product, error) {

	if limit < 0 {
		return nil, utility.InvalidQuantity
	}

	return s.repository.FindLowStock(limit)
}
