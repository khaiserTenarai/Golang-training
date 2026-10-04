package controller

import (
	"fmt"

	"prodcut-inventory/service"
	"prodcut-inventory/view"
)

type ProductController struct {
	service service.ProductService
	view    *view.ProductView
}

func NewProductController(
	service service.ProductService,
	view *view.ProductView,
) *ProductController {

	return &ProductController{
		service: service,
		view:    view,
	}
}

func (c *ProductController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.AddProduct()

		case 2:
			c.FindProductByID()

		case 3:
			c.FindAllProducts()

		case 4:
			c.UpdateProduct()

		case 5:
			c.DeleteProduct()

		case 6:
			c.IncreaseStock()

		case 7:
			c.DecreaseStock()

		case 8:
			c.FindLowStock()

		case 9:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// CREATE
func (c *ProductController) AddProduct() {

	product := c.view.ReadProduct()

	err := c.service.AddProduct(product)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Product added successfully.",
	)
}

// READ ONE
func (c *ProductController) FindProductByID() {

	id := c.view.ReadInt(
		"Enter Product ID: ",
	)

	product, err :=
		c.service.FindProductByID(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowProduct(product)
}

// READ ALL
func (c *ProductController) FindAllProducts() {

	products :=
		c.service.FindAllProducts()

	c.view.ShowProducts(products)
}

// UPDATE
func (c *ProductController) UpdateProduct() {

	product :=
		c.view.ReadProductForUpdate()

	err :=
		c.service.UpdateProduct(product)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Product updated successfully.",
	)
}

// DELETE
func (c *ProductController) DeleteProduct() {

	id := c.view.ReadInt(
		"Enter Product ID: ",
	)

	err :=
		c.service.DeleteProduct(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Product deleted successfully.",
	)
}

// INCREASE STOCK
func (c *ProductController) IncreaseStock() {

	id := c.view.ReadInt(
		"Enter Product ID: ",
	)

	quantity := c.view.ReadInt(
		"Enter Quantity: ",
	)

	err :=
		c.service.IncreaseStock(
			id,
			quantity,
		)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Stock increased successfully.",
	)
}

// DECREASE STOCK
func (c *ProductController) DecreaseStock() {

	id := c.view.ReadInt(
		"Enter Product ID: ",
	)

	quantity := c.view.ReadInt(
		"Enter Quantity: ",
	)

	err :=
		c.service.DecreaseStock(
			id,
			quantity,
		)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Stock decreased successfully.",
	)
}

// LOW STOCK
func (c *ProductController) FindLowStock() {

	products :=
		c.service.FindLowStock()

	c.view.ShowProducts(products)
}
