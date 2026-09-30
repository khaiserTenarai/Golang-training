package utility
import (
	"errors"
	"strings"

	"product_inventory/model"
)

var (
	InvalidID       = errors.New("invalid product ID")
	InvalidName     = errors.New("invalid product name")
	InvalidPrice    = errors.New("invalid product price")
	InvalidQuantity = errors.New("invalid quantity")
)

func ValidateProduct(product model.Product) error {

	if strings.TrimSpace(product.Name) == "" {
		return InvalidName
	}

	if product.Price <= 0 {
		return InvalidPrice
	}

	if product.Quantity < 0 {
		return InvalidQuantity
	}

	return nil
}

func ValidateID(id int) error {

	if id <= 0 {
		return InvalidID
	}

	return nil
}

func ValidateQuantity(quantity int) error {

	if quantity <= 0 {
		return InvalidQuantity
	}

	return nil
}
