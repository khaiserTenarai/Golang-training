package service

import (
	"fmt"

	"example.com/employee_management/model"
	"example.com/employee_management/repository"
)

type EmployeeService interface {
	AddEmployee(employee model.Employee)
	GetEmployee(id int)
	GetAllEmployees()
	UpdateEmployee(employee model.Employee)
	DeleteEmployee(id int)
}

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{
		repository: employeeRepository,
	}
}

func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) {
	s.repository.AddEmployee(employee)
	fmt.Println("Employee added successfully")
}

func (s *EmployeeServiceImpl) GetEmployee(id int) {
	employee, err := s.repository.GetEmployee(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	employee.Display()
}

func (s *EmployeeServiceImpl) GetAllEmployees() {
	employees := s.repository.GetAllEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}

	for _, employee := range employees {
		fmt.Println("----------------------------")
		employee.Display()
	}
}

func (s *EmployeeServiceImpl) UpdateEmployee(employee model.Employee) {
	err := s.repository.UpdateEmployee(employee)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Employee updated successfully")
}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) {
	err := s.repository.DeleteEmployee(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Employee deleted successfully")
}