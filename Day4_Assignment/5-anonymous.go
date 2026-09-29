package main
import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	employees := []Employee{
		{"Vittesh", 30000},
		{"dhanu", 50000},
		{"guru", 25000},
	}

	filterSalary := func(employee Employee) bool {
		return employee.Salary > 30000
	}

	for _, employee := range employees {

		if filterSalary(employee) {
			fmt.Println(employee.Name, employee.Salary)
		}
	}
}