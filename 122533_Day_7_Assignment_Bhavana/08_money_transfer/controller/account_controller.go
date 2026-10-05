package controller

import (
	"bufio"
	"context"
	"example.com/q8-money-transfer/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AccountController struct {
	service service.AccountService
	reader  *bufio.Reader
}

func NewAccountController(s service.AccountService) *AccountController {
	return &AccountController{service: s, reader: bufio.NewReader(os.Stdin)}
}
func (c *AccountController) Start() {
	for {
		fmt.Println("\n1. Create Account\n2. View Accounts\n3. Transfer Money\n4. Exit")
		choice := c.readInt("Choice: ")
		ctx := context.Background()
		switch choice {
		case 1:
			name := c.readString("Name: ")
			balance := c.readFloat("Opening balance: ")
			c.show(c.service.CreateAccount(ctx, name, balance), "Account created.")
		case 2:
			c.show(c.service.ListAccounts(ctx), "")
		case 3:
			from := c.readInt64("From account ID: ")
			to := c.readInt64("To account ID: ")
			amount := c.readFloat("Amount: ")
			c.show(c.service.Transfer(ctx, from, to, amount), "Transfer completed.")
		case 4:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (c *AccountController) show(err error, success string) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if success != "" {
		fmt.Println(success)
	}
}
func (c *AccountController) readString(p string) string {
	fmt.Print(p)
	v, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(v)
}
func (c *AccountController) readInt(p string) int { v, _ := strconv.Atoi(c.readString(p)); return v }
func (c *AccountController) readInt64(p string) int64 {
	v, _ := strconv.ParseInt(c.readString(p), 10, 64)
	return v
}
func (c *AccountController) readFloat(p string) float64 {
	v, _ := strconv.ParseFloat(c.readString(p), 64)
	return v
}
