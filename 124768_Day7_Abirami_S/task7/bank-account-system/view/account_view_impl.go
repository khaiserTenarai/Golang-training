package view

import (
	"bank-account-system/controller"
	"bank-account-system/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AccountViewImpl struct {
	controller controller.AccountController
	reader     *bufio.Reader
}

func NewAccountView(controller controller.AccountController) AccountView {
	return &AccountViewImpl{
		controller: controller,
		reader:     bufio.NewReader(os.Stdin),
	}
}

func (v *AccountViewImpl) Start() {
	for {
		fmt.Println("\n----- BANK ACCOUNT SYSTEM -----")
		fmt.Println("1. Create Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Balance Enquiry")
		fmt.Println("5. Transaction History")
		fmt.Println("6. Exit")

		fmt.Print("Enter choice: ")
		choice := v.readInt()

		switch choice {
		case 1:
			v.CreateAccount()
		case 2:
			v.Deposit()
		case 3:
			v.Withdraw()
		case 4:
			v.BalanceEnquiry()
		case 5:
			v.TransactionHistory()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (v *AccountViewImpl) CreateAccount() {
	var account model.Account

	fmt.Print("Enter Account Number: ")
	account.AccountNumber = v.readString()

	fmt.Print("Enter Holder Name: ")
	account.HolderName = v.readString()

	fmt.Print("Enter Initial Balance: ")
	account.Balance = v.readFloat()

	err := v.controller.CreateAccount(account)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Account created successfully")
}

func (v *AccountViewImpl) Deposit() {
	fmt.Print("Enter Account ID: ")
	id := v.readInt()

	fmt.Print("Enter Deposit Amount: ")
	amount := v.readFloat()

	err := v.controller.Deposit(id, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Deposit successful")
}

func (v *AccountViewImpl) Withdraw() {
	fmt.Print("Enter Account ID: ")
	id := v.readInt()

	fmt.Print("Enter Withdrawal Amount: ")
	amount := v.readFloat()

	err := v.controller.Withdraw(id, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Withdrawal successful")
}

func (v *AccountViewImpl) BalanceEnquiry() {
	fmt.Print("Enter Account ID: ")
	id := v.readInt()

	account, err := v.controller.GetAccount(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayAccount(account)
}

func (v *AccountViewImpl) TransactionHistory() {
	fmt.Print("Enter Account ID: ")
	accountID := v.readInt()

	transactions, err := v.controller.GetTransactionHistory(accountID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayTransactions(transactions)
}

func (v *AccountViewImpl) DisplayAccount(account *model.Account) {
	fmt.Println("\n----- ACCOUNT DETAILS -----")
	fmt.Println("ID             :", account.ID)
	fmt.Println("Account Number :", account.AccountNumber)
	fmt.Println("Holder Name    :", account.HolderName)
	fmt.Println("Balance        :", account.Balance)
}

func (v *AccountViewImpl) DisplayTransactions(transactions []model.Transaction) {
	fmt.Println("\n----- TRANSACTION HISTORY -----")

	if len(transactions) == 0 {
		fmt.Println("No transactions found")
		return
	}

	for _, transaction := range transactions {
		fmt.Println("----------------------------")
		fmt.Println("ID              :", transaction.ID)
		fmt.Println("Account ID      :", transaction.AccountID)
		fmt.Println("Type            :", transaction.TransactionType)
		fmt.Println("Amount          :", transaction.Amount)
		fmt.Println("Transaction Date:", transaction.TransactionDate)
	}
}

func (v *AccountViewImpl) readString() string {
	value, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func (v *AccountViewImpl) readInt() int {
	value := v.readString()
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return number
}

func (v *AccountViewImpl) readFloat() float64 {
	value := v.readString()
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return number
}
