package controller

import (
	"fmt"

	"money-transfer/service"
	"money-transfer/view"
)

type AccountController struct {
	service service.AccountService
	view    *view.AccountView
}

func NewAccountController(
	service service.AccountService,
	view *view.AccountView,
) *AccountController {

	return &AccountController{
		service: service,
		view:    view,
	}
}

func (c *AccountController) Start() {

	for {

		c.view.ShowMenu()

		choice := c.view.ReadInt(
			"Enter your choice: ",
		)

		switch choice {

		case 1:
			c.TransferMoney()

		case 2:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *AccountController) TransferMoney() {

	fromID := c.view.ReadInt(
		"Enter Sender Account ID: ",
	)

	toID := c.view.ReadInt(
		"Enter Receiver Account ID: ",
	)

	amount := c.view.ReadFloat(
		"Enter Amount: ",
	)

	err := c.service.Transfer(
		fromID,
		toID,
		amount,
	)

	if err != nil {

		c.view.ShowMessage(
			"Error: " + err.Error(),
		)

		return
	}

	c.view.ShowMessage(
		"Money transferred successfully.",
	)
}
