package main

import "fmt"

func createEmployee(id int, name string, salary float64) {
	fmt.Println("Employee Created")
	fmt.Println("ID:", id)
	fmt.Println("NAME:", name)
	fmt.Println("SALARY:", salary)
}

func readEmployee(id int, name string, salary float64) {
	fmt.Println("\nEmployee Details:")
	fmt.Println("ID:", id)
	fmt.Println("NAME:", name)
	fmt.Println("SALARY:", salary)
}
func updateEmployee(name string, salary float64) (string, float64) {
	name = "Muneera Shaik"
	salary = 30000.0
	fmt.Println("Employee Details updated")
	return name, salary

}

func deleteEmployee(id int, name string, salary float64) {
	id = 0
	name = ""
	salary = 0
	fmt.Println("\nEmployee Deleted:")
	fmt.Println("ID:", id)
	fmt.Println("NAME:", name)
	fmt.Println("SALARY:", salary)
}

func main() {
	id := 101
	name := "Muneera"
	salary := 20000.0

	createEmployee(id, name, salary)
	readEmployee(id, name, salary)
	name, salary = updateEmployee(name, salary)
	fmt.Println("\nAfter Employee Update:")
	readEmployee(id, name, salary)

	deleteEmployee(id, name, salary)
}
