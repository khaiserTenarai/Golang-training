package view
import (
	"fmt"

	"bank_account/model"
)

type AccountViewImpl struct {
}

func NewAccountView() AccountView {

	return &AccountViewImpl{}
}

func (v *AccountViewImpl) ShowMenu() int {

	fmt.Println("\n========== Bank Account System ==========")

	fmt.Println("1. Create Account")
	fmt.Println("2. Balance Enquiry")
	fmt.Println("3. Deposit")
	fmt.Println("4. Withdraw")
	fmt.Println("5. Transaction History")
	fmt.Println("6. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *AccountViewImpl) ReadAccount() model.Account {

	var account model.Account

	fmt.Println("\n---------- Create Account ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&account.Name)

	fmt.Print("Enter Email: ")
	fmt.Scan(&account.Email)

	fmt.Print("Enter Initial Balance: ")
	fmt.Scan(&account.Balance)

	return account
}

func (v *AccountViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Account ID: ")
	fmt.Scan(&id)

	return id
}

func (v *AccountViewImpl) ReadAmount() float64 {

	var amount float64

	fmt.Print("Enter Amount: ")
	fmt.Scan(&amount)

	return amount
}

func (v *AccountViewImpl) DisplayAccount(
	account model.Account,
) {

	fmt.Println("\n---------- Account ----------")

	fmt.Println("ID      :", account.ID)
	fmt.Println("Name    :", account.Name)
	fmt.Println("Email   :", account.Email)
	fmt.Println("Balance :", account.Balance)
}

func (v *AccountViewImpl) DisplayTransactions(
	transactions []model.Transaction,
) {

	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Println("\n---------- Transaction History ----------")

	for _, transaction := range transactions {

		fmt.Println(
			"Transaction ID:",
			transaction.ID,
			"Type:",
			transaction.Type,
			"Amount:",
			transaction.Amount,
		)
	}
}
