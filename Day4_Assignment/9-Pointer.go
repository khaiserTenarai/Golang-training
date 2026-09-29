package main
import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func increaseSalary(employee *Employee) {

	employee.Salary = employee.Salary + 5000
}

func main() {

	employee := Employee{
		Name:   "Vittesh",
		Salary: 30000,
	}

	fmt.Println("Before:", employee.Salary)

	increaseSalary(&employee)

	fmt.Println("After:", employee.Salary)
}