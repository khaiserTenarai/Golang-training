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
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Address    Address
	Department Department
}


func (e Employee) Describe() string {
	return fmt.Sprintf("[%d] %s works in %s, earns %.2f", e.ID, e.Name, e.Department.Name, e.Salary)
}

func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}


func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
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

	fmt.Println(e.Describe())
	fmt.Println("Annual salary:", e.AnnualSalary())

	e.GiveRaise(10)
	fmt.Println("After a 10% raise:", e.Describe())
}
