package view
import "product_inventory/model"

type ProductView interface {

	ShowMenu() int

	ReadProduct() model.Product

	ReadProductForUpdate() model.Product

	ReadID() int

	ReadStockQuantity() int

	ReadLowStockLimit() int

	DisplayProduct(product model.Product)

	DisplayProducts(products []model.Product)
}
