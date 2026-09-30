package controller
import (
	"fmt"

	"product_inventory/service"
	"product_inventory/view"
)

type ProductControllerImpl struct {
	view    view.ProductView
	service service.ProductService
}

func NewProductController(
	view view.ProductView,
	service service.ProductService,
) ProductController {

	return &ProductControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *ProductControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 9 {
			fmt.Println("Thank you..")
			return
		}

		c.Process(choice)
	}
}

func (c *ProductControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		product := c.view.ReadProduct()

		err := c.service.Save(product)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Product saved successfully.")

	case 2:

		id := c.view.ReadID()

		product, err := c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayProduct(product)

	case 3:

		products, err := c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayProducts(products)

	case 4:

		product := c.view.ReadProductForUpdate()

		err := c.service.Update(product)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Product updated successfully.")

	case 5:

		id := c.view.ReadID()

		err := c.service.Delete(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Product deleted successfully.")

	case 6:

		id := c.view.ReadID()

		quantity := c.view.ReadStockQuantity()

		err := c.service.IncreaseStock(
			id,
			quantity,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Stock increased successfully.")

	case 7:

		id := c.view.ReadID()

		quantity := c.view.ReadStockQuantity()

		err := c.service.DecreaseStock(
			id,
			quantity,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Stock decreased successfully.")

	case 8:

		limit := c.view.ReadLowStockLimit()

		products, err :=
			c.service.FindLowStock(limit)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayProducts(products)

	default:

		fmt.Println("Invalid choice.")
	}
}
