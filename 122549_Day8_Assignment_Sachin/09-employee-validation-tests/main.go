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

	fmt.Print("Enter employee name: ")
	nameText, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameText)

	fmt.Print("Enter employee email: ")
	emailText, _ := reader.ReadString('\n')
	email := strings.TrimSpace(emailText)

	fmt.Print("Enter employee age: ")
	ageText, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageText))
	if err != nil {
		fmt.Println("Age must be a number.")
		return
	}

	if err := validateName(name); err != nil {
		fmt.Println("Name invalid:", err)
	}
	if err := validateEmail(email); err != nil {
		fmt.Println("Email invalid:", err)
	}
	if err := validateAge(age); err != nil {
		fmt.Println("Age invalid:", err)
	}

	fmt.Println("Validation complete.")
}
