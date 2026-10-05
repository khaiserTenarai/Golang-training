package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter purchase amount: ")
	amountText, _ := reader.ReadString('\n')
	amount, err := strconv.ParseFloat(strings.TrimSpace(amountText), 64)
	if err != nil {
		fmt.Println("Please enter a valid number.")
		return
	}

	fmt.Printf("Discount: %.0f%%\n", DiscountPercent(amount))
}
