package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strconv.ParseFloat(strings.TrimSpace(text), 64)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	basic, err := readFloat(reader, "Enter basic salary: ")
	if err != nil {
		fmt.Println("Invalid basic salary.")
		return
	}

	allowance, err := readFloat(reader, "Enter allowance: ")
	if err != nil {
		fmt.Println("Invalid allowance.")
		return
	}

	deduction, err := readFloat(reader, "Enter deduction: ")
	if err != nil {
		fmt.Println("Invalid deduction.")
		return
	}

	fmt.Println("Net Salary:", netSalary(basic, allowance, deduction))
}
