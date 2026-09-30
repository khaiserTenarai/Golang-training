package controller
import (
	"fmt"

	"ecommerce/service"
	"ecommerce/view"
)

type OrderController interface {
	Start()
	Process(choice int)
}

type OrderControllerImpl struct {
	view    view.OrderView
	service service.OrderService
}

func NewOrderController(
	view view.OrderView,
	service service.OrderService,
) OrderController {

	return &OrderControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *OrderControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 4 {
			return
		}

		c.Process(choice)
	}
}

func (c *OrderControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		order := c.view.ReadOrder()

		err := c.service.CreateOrder(order)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Order created successfully.")

	case 2:

		id := c.view.ReadID()

		details, err := c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayOrder(details)

	case 3:

		details, err := c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayOrders(details)

	default:

		fmt.Println("Invalid choice.")
	}
}
