
package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}


func modifyByValue(e Employee) {
	e.Salary = 99999
}

func modifyByPointer(e *Employee) {
	e.Salary = 99999
}

func main() {
	ranjitha := Employee{Name: "Ranjitha", Salary: 60000}

	modifyByValue(ranjitha)
	fmt.Println("After modifyByValue  :", ranjitha.Salary) // unchanged

	modifyByPointer(&ranjitha)
	fmt.Println("After modifyByPointer:", ranjitha.Salary) // changed
}
