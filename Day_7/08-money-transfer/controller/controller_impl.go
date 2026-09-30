package controller
import (
	"fmt"

	"money_transfer/service"
	"money_transfer/view"
)

type TransferControllerImpl struct {
	view    view.TransferView
	service service.TransferService
}

func NewTransferController(
	view view.TransferView,
	service service.TransferService,
) TransferController {

	return &TransferControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *TransferControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 2 {

			fmt.Println("Thank you..")
			return
		}

		c.Process(choice)
	}
}

func (c *TransferControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		transfer := c.view.ReadTransfer()

		err := c.service.TransferMoney(
			transfer,
		)

		if err != nil {

			fmt.Println("Transfer failed:", err)
			return
		}

		fmt.Println(
			"Money transferred successfully.",
		)

	default:

		fmt.Println("Invalid choice.")
	}
}
