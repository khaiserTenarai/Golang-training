package controller

import (
	"fmt"

	"ecommerce-order/model"
	"ecommerce-order/service"
	"ecommerce-order/view"
)

type OrderController struct {
	service service.OrderService
	view    *view.OrderView
}

func NewOrderController(
	service service.OrderService,
	view *view.OrderView,
) *OrderController {

	return &OrderController{
		service: service,
		view:    view,
	}
}

func (c *OrderController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.CreateCustomer()

		case 2:
			c.CreateProduct()

		case 3:
			c.CreateOrder()

		case 4:
			c.AddOrderItem()

		case 5:
			c.ViewOrder()

		case 6:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *OrderController) CreateCustomer() {

	customer := c.view.ReadCustomer()

	err := c.service.CreateCustomer(customer)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowMessage(
		"Customer created successfully.",
	)
}

func (c *OrderController) CreateProduct() {

	product := c.view.ReadProduct()

	err := c.service.CreateProduct(product)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowMessage(
		"Product created successfully.",
	)
}

func (c *OrderController) CreateOrder() {

	customerID := c.view.ReadInt(
		"Enter Customer ID: ",
	)

	orderID, err :=
		c.service.CreateOrder(customerID)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	fmt.Println(
		"Order created successfully. Order ID:",
		orderID,
	)
}

func (c *OrderController) AddOrderItem() {

	item := model.OrderItem{}

	item.OrderID = c.view.ReadInt(
		"Enter Order ID: ",
	)

	item.ProductID = c.view.ReadInt(
		"Enter Product ID: ",
	)

	item.Quantity = c.view.ReadInt(
		"Enter Quantity: ",
	)

	item.Price = c.view.ReadFloat(
		"Enter Price: ",
	)

	err := c.service.AddOrderItem(item)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowMessage(
		"Order item added successfully.",
	)
}

func (c *OrderController) ViewOrder() {

	orderID := c.view.ReadInt(
		"Enter Order ID: ",
	)

	details, err :=
		c.service.GetOrderDetails(orderID)

	if err != nil {
		c.view.ShowMessage("Error: " + err.Error())
		return
	}

	c.view.ShowOrderDetails(details)
}
