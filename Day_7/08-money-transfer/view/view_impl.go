package view
import (
	"fmt"

	"money_transfer/model"
)

type TransferViewImpl struct {
}

func NewTransferView() TransferView {

	return &TransferViewImpl{}
}

func (v *TransferViewImpl) ShowMenu() int {

	fmt.Println("\n========== Money Transfer System ==========")

	fmt.Println("1. Transfer Money")
	fmt.Println("2. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *TransferViewImpl) ReadTransfer() model.Transfer {

	var transfer model.Transfer

	fmt.Println("\n---------- Money Transfer ----------")

	fmt.Print("Enter Sender Account ID: ")
	fmt.Scan(&transfer.FromAccountID)

	fmt.Print("Enter Receiver Account ID: ")
	fmt.Scan(&transfer.ToAccountID)

	fmt.Print("Enter Amount: ")
	fmt.Scan(&transfer.Amount)

	return transfer
}
