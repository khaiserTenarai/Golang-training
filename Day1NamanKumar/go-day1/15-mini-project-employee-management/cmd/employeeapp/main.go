// Command employeeapp is an interactive Employee Management CLI.
//
// Data is kept in memory for the lifetime of the program. Set
// SEED_DATA=false to start with an empty list.
package main

import (
	"fmt"
	"os"

	"employeeapp/internal/cli"
	"employeeapp/internal/employee"
)

func main() {
	store := employee.NewStore()
	if os.Getenv("SEED_DATA") != "false" {
		seed(store)
	}
	fmt.Println("Welcome to the Employee Management CLI")
	cli.New(store, os.Stdin, os.Stdout).Run()
}

func seed(s *employee.Store) {
	for _, e := range []employee.Employee{
		{Name: "Asha Rao", Department: "Engineering", Salary: 1200000},
		{Name: "Vikram Iyer", Department: "Finance", Salary: 950000},
		{Name: "Meera Nair", Department: "HR", Salary: 800000},
	} {
		if _, err := s.Add(e); err != nil {
			fmt.Fprintln(os.Stderr, "seed error:", err)
		}
	}
}
