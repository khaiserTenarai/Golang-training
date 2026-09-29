package main
import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

// Value
func changeSalaryValue(employee Employee) {

	employee.Salary = 50000
}

// Pointer
func changeSalaryPointer(employee *Employee) {

	employee.Salary = 50000
}

func main() {

	employee := Employee{
		Name:   "Vittesh",
		Salary: 30000,
	}

	changeSalaryValue(employee)

	fmt.Println("After value:", employee.Salary)

	changeSalaryPointer(&employee)

	fmt.Println("After pointer:", employee.Salary)
}