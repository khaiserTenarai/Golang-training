package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OrderView struct {
	reader *bufio.Reader
}

func NewOrderView() *OrderView {

	return &OrderView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *OrderView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== ORDER & INVENTORY ==========")
	fmt.Println("1. Create Order")
	fmt.Println("2. Exit")
	fmt.Println("=======================================")
}

func (v *OrderView) ReadString(
	message string,
) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (v *OrderView) ReadInt(
	message string,
) int {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Enter a valid number.")
	}
}

func (v *OrderView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}
