package main

import "fmt"

type Employee9 struct {
	ID   int
	Name string
}
type EmployeeRepository interface {
	GetEmployee(id int) Employee9
	SaveEmployee(employee Employee9)
}
type EmployeeRepositoryImpl struct{}

func (r EmployeeRepositoryImpl) GetEmployee(id int) Employee9 {
	return Employee9{
		ID:   id,
		Name: "Abirami",
	}
}
func (r EmployeeRepositoryImpl) SaveEmployee(employee Employee9) {
	fmt.Println("Employee saved: ", employee)
}
func main() {
	var repository EmployeeRepository = EmployeeRepositoryImpl{}
	employee := repository.GetEmployee(101)
	fmt.Println("Employee: ", employee)
	repository.SaveEmployee(employee)
}
