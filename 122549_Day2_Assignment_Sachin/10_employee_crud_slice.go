package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func main() {
	list := []Employee{
		{1, "Sachin"},
		{2, "nakul"},
		{3, "bhavana"},
	}

	fmt.Println("**Initial List**")
	for _, e := range list {
		fmt.Println(e.ID, e.Name)
	}

	for i := range list {
		if list[i].ID == 2 {
			list[i].Name = "Sneha Patil"
		}
	}

	var updated []Employee
	for _, e := range list {
		if e.ID != 1 {
			updated = append(updated, e)
		}
	}
	list = updated

	fmt.Println("-- After Update & Delete --")
	for _, e := range list {
		fmt.Println(e.ID, e.Name)
	}
}