package main

import "fmt"

type Address struct {
	City    string `json:"city"`
	Pincode string `json:"pincode"`
}

type Department struct {
	Name string `json:"name"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func main() {
	e := Employee{
		ID:      1,
		Name:    "Ray",
		Email:   "ray@example.com",
		Age:     26,
		Salary:  60000,
		Address: Address{City: "Bengaluru", Pincode: "560001"},
		Department: Department{
			Name: "Engineering",
		},
	}

	fmt.Println("Struct fields now carry JSON tags (see source code).")
	fmt.Println("Go field 'Name' will serialize as JSON key \"name\", etc.")
	fmt.Printf("%+v\n", e)
	fmt.Println("\nSee step4 to actually convert this to JSON text.")
}
