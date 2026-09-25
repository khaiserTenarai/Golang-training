package main

import (
	"company/employee"
)

func main() {
	emp := employee.New(101, "John Doe", 50000)

	emp.Display()

	emp = emp.GiveRaise(5000)

	emp.Display()
}
