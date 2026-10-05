package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	service := NewEmployeeService()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("---- Employee Service ----")
		fmt.Println("1. Add employee")
		fmt.Println("2. Get employee")
		fmt.Println("3. Delete employee")
		fmt.Println("4. Count employees")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		choiceText, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceText)

		switch choice {
		case "1":
			fmt.Print("Name: ")
			nameText, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameText)

			fmt.Print("Salary: ")
			salaryText, _ := reader.ReadString('\n')
			salary, err := strconv.ParseFloat(strings.TrimSpace(salaryText), 64)
			if err != nil {
				fmt.Println("Salary must be a number.")
				continue
			}

			id := service.Add(name, salary)
			fmt.Println("Added with ID:", id)

		case "2":
			fmt.Print("ID: ")
			idText, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idText))
			if err != nil {
				fmt.Println("ID must be a number.")
				continue
			}
			emp, err := service.Get(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%+v\n", emp)

		case "3":
			fmt.Print("ID: ")
			idText, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idText))
			if err != nil {
				fmt.Println("ID must be a number.")
				continue
			}
			if err := service.Delete(id); err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Deleted.")

		case "4":
			fmt.Println("Total employees:", service.Count())

		case "5":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid option, please choose 1-5.")
		}
	}
}
