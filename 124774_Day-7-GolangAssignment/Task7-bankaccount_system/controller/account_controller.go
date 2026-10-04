package controller

import (
	"fmt"

	"bankaccount/service"
	"bankaccount/view"
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
			c.CreateAccount()

		case 2:
			c.Deposit()

		case 3:
			c.Withdraw()

		case 4:
			c.FindAccount()

		case 5:
			c.FindTransactions()

		case 6:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// CREATE ACCOUNT

func (c *AccountController) CreateAccount() {

	account := c.view.ReadAccount()

	err := c.service.CreateAccount(account)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Account created successfully.",
	)
}

// DEPOSIT

func (c *AccountController) Deposit() {

	id := c.view.ReadInt(
		"Enter Account ID: ",
	)

	amount := c.view.ReadFloat(
		"Enter Deposit Amount: ",
	)

	err := c.service.Deposit(
		id,
		amount,
	)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Deposit successful.",
	)
}

// WITHDRAW

func (c *AccountController) Withdraw() {

	id := c.view.ReadInt(
		"Enter Account ID: ",
	)

	amount := c.view.ReadFloat(
		"Enter Withdrawal Amount: ",
	)

	err := c.service.Withdraw(
		id,
		amount,
	)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowMessage(
		"Withdrawal successful.",
	)
}

// BALANCE ENQUIRY

func (c *AccountController) FindAccount() {

	id := c.view.ReadInt(
		"Enter Account ID: ",
	)

	account, err := c.service.FindAccount(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowAccount(account)
}

// TRANSACTION HISTORY

func (c *AccountController) FindTransactions() {

	id := c.view.ReadInt(
		"Enter Account ID: ",
	)

	transactions, err :=
		c.service.FindTransactions(id)

	if err != nil {
		c.view.ShowMessage(
			"Error: " + err.Error(),
		)
		return
	}

	c.view.ShowTransactions(transactions)
}
