package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	employees := []Employee{
		{ID: 1, Name: "Ray", Salary: 50000},
		{ID: 2, Name: "John", Salary: 60000},
		{ID: 3, Name: "Sam", Salary: 55000},
	}

	fmt.Println("Monthly Employee Report:")

	for _, employee := range employees {
		fmt.Println(employee)

	}
	fmt.Println("Employee Management Application")
}

// my_personal_git_account_link = https://github.com/RanjithaSelin/employee-management-git/blob/main/Prac

