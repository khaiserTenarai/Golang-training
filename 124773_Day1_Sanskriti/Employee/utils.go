package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Global reader to handle inputs with spaces safely
var reader = bufio.NewReader(os.Stdin)

func PrintDivider() {
	fmt.Println("------------------------------")
}

func ShowMenu() {
	PrintDivider()
	fmt.Println(" EMPLOYEE MANAGEMENT SYSTEM")
	PrintDivider()
	fmt.Println("1. Add Employee")
	fmt.Println("2. Search Employee")
	fmt.Println("3. Display All Employees")
	fmt.Println("4. Delete Employee")
	fmt.Println("5. Exit")
	PrintDivider()
}

// Helper to read single line input from user
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}