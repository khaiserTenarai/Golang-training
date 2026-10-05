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

	fmt.Print("Enter a number: ")
	text, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil || n < 0 {
		fmt.Println("Please enter a non-negative whole number.")
		return
	}

	fmt.Println("Factorial:", factorial(n))
}
