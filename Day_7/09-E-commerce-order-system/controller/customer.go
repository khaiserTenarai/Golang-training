package controller
import (
	"fmt"

	"ecommerce/service"
	"ecommerce/view"
)

type CustomerController interface {
	Start()
	Process(choice int)
}

type CustomerControllerImpl struct {
	view    view.CustomerView
	service service.CustomerService
}

func NewCustomerController(
	view view.CustomerView,
	service service.CustomerService,
) CustomerController {

	return &CustomerControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *CustomerControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 6 {
			return
		}

		c.Process(choice)
	}
}

func (c *CustomerControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		customer := c.view.ReadCustomer()

		err := c.service.Save(customer)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Customer saved successfully.")

	case 2:

		id := c.view.ReadID()

		customer, err := c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayCustomer(customer)

	case 3:

		customers, err := c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayCustomers(customers)

	case 4:

		customer := c.view.ReadCustomerForUpdate()

		err := c.service.Update(customer)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Customer updated successfully.")

	case 5:

		id := c.view.ReadID()

		err := c.service.Delete(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Customer deleted successfully.")

	default:

		fmt.Println("Invalid choice.")
	}
}
