package controller

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "money-transfer-system/model"
    "money-transfer-system/service"
)

type TransferController struct {
    transferService service.TransferService
    reader          *bufio.Reader
}

func NewTransferController(
    transferService service.TransferService,
) *TransferController {
    return &TransferController{
        transferService: transferService,
        reader:          bufio.NewReader(os.Stdin),
    }
}

func (c *TransferController) Start() {
    for {
        fmt.Println()
        fmt.Println("======================================")
        fmt.Println("         MONEY TRANSFER SYSTEM")
        fmt.Println("======================================")
        fmt.Println("1. Transfer Money")
        fmt.Println("2. Check Account")
        fmt.Println("3. Exit")
        fmt.Println("======================================")

        choice := c.readInt("Enter choice: ")

        switch choice {
        case 1:
            c.transferMoney()
        case 2:
            c.checkAccount()
        case 3:
            fmt.Println("Thank you for using Money Transfer System.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func (c *TransferController) transferMoney() {
    transfer := model.Transfer{
        FromAccount: c.readString("Enter source account: "),
        ToAccount:   c.readString("Enter destination account: "),
        Amount:      c.readFloat("Enter amount: "),
    }

    fmt.Println()
    fmt.Println("Processing transfer...")

    err := c.transferService.TransferMoney(transfer)
    if err != nil {
        fmt.Println("Transfer failed:", err)
        fmt.Println("Database transaction has been rolled back.")
        return
    }

    fmt.Println("Transfer successful.")
    fmt.Println("Transaction committed successfully.")
}

func (c *TransferController) checkAccount() {
    accountNumber := c.readString("Enter account number: ")

    account, err := c.transferService.GetAccount(accountNumber)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("----------- ACCOUNT -----------")
    fmt.Println("Account Number:", account.AccountNumber)
    fmt.Println("Customer Name:", account.CustomerName)
    fmt.Printf("Balance: ₹%.2f
", account.Balance)
    fmt.Println("-------------------------------")
}

func (c *TransferController) readString(prompt string) string {
    fmt.Print(prompt)
    value, _ := c.reader.ReadString('
')
    return strings.TrimSpace(value)
}

func (c *TransferController) readInt(prompt string) int {
    for {
        value := c.readString(prompt)

        number, err := strconv.Atoi(value)
        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid number.")
    }
}

func (c *TransferController) readFloat(prompt string) float64 {
    for {
        value := c.readString(prompt)

        number, err := strconv.ParseFloat(value, 64)
        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid amount.")
    }
}
