package utility
import (
	"errors"
	"strings"

	"order_inventory/model"
)

var (
	InvalidIDError       = errors.New("invalid ID")
	InvalidNameError     = errors.New("invalid name")
	InvalidPriceError    = errors.New("price must be greater than 0")
	InvalidStockError    = errors.New("stock cannot be negative")
	InvalidQuantityError = errors.New("quantity must be greater than 0")
	InvalidOrderError    = errors.New("order must contain at least one item")
)

func ValidateID(id int) error {

	if id <= 0 {
		return InvalidIDError
	}

	return nil
}

func ValidateProduct(
	product model.Product,
) error {

	if strings.TrimSpace(product.Name) == "" {
		return InvalidNameError
	}

	if product.Price <= 0 {
		return InvalidPriceError
	}

	if product.Stock < 0 {
		return InvalidStockError
	}

	return nil
}

func ValidateOrder(
	order model.OrderRequest,
) error {

	if strings.TrimSpace(order.CustomerName) == "" {
		return InvalidNameError
	}

	if len(order.Items) == 0 {
		return InvalidOrderError
	}

	for _, item := range order.Items {

		if item.ProductID <= 0 {
			return InvalidIDError
		}

		if item.Quantity <= 0 {
			return InvalidQuantityError
		}
	}

	return nil
}
