package controller
import (
	"fmt"

	"bank_account/service"
	"bank_account/view"
)

type AccountControllerImpl struct {
	view    view.AccountView
	service service.AccountService
}

func NewAccountController(
	view view.AccountView,
	service service.AccountService,
) AccountController {

	return &AccountControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *AccountControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 6 {
			fmt.Println("Thank you..")
			return
		}

		c.Process(choice)
	}
}

func (c *AccountControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		account := c.view.ReadAccount()

		err := c.service.Create(account)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Account created successfully.")

	case 2:

		id := c.view.ReadID()

		account, err := c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayAccount(account)

	case 3:

		id := c.view.ReadID()

		amount := c.view.ReadAmount()

		err := c.service.Deposit(
			id,
			amount,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Amount deposited successfully.")

	case 4:

		id := c.view.ReadID()

		amount := c.view.ReadAmount()

		err := c.service.Withdraw(
			id,
			amount,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Amount withdrawn successfully.")

	case 5:

		id := c.view.ReadID()

		transactions, err :=
			c.service.GetTransactions(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayTransactions(
			transactions,
		)

	default:

		fmt.Println("Invalid choice.")
	}
}
