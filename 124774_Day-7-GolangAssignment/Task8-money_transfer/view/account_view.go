package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AccountView struct {
	reader *bufio.Reader
}

func NewAccountView() *AccountView {

	return &AccountView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *AccountView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== MONEY TRANSFER ==========")
	fmt.Println("1. Transfer Money")
	fmt.Println("2. Exit")
	fmt.Println("====================================")
}

func (v *AccountView) ReadInt(message string) int {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid number.")
	}
}

func (v *AccountView) ReadFloat(message string) float64 {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid amount.")
	}
}

func (v *AccountView) ShowMessage(message string) {

	fmt.Println(message)
}
