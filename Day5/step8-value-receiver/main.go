package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) IsSenior(ageThreshold int, age int) bool {
	return age >= ageThreshold
}

func (e Employee) AttemptRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
	fmt.Println("Inside AttemptRaise (value receiver), salary is now:", e.Salary)
}

func main() {
	e := Employee{Name: "Ray", Salary: 60000}

	fmt.Println("Before AttemptRaise:", e.Salary)
	e.AttemptRaise(10)
	fmt.Println("After AttemptRaise: ", e.Salary, "<- unchanged! (value receiver got a copy)")

	fmt.Println("\nCompare this with step7's GiveRaise, which uses a")
	fmt.Println("pointer receiver and DOES change the original.")
}
