package controller
import (
	"fmt"

	"order_inventory/service"
	"order_inventory/view"
)

type ProductController interface {
	Start()
	Process(choice int)
}

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

		if choice == 5 {
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

	default:

		fmt.Println("Invalid choice.")
	}
}
