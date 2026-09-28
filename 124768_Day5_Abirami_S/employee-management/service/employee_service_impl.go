package service

import (
	"employee-management/model"
	"employee-management/repository"
	"employee-management/utility"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(repository repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{
		repository: repository,
	}
}
func (s *EmployeeServiceImpl) AddEmployee(employee model.Employee) error {
	err := utility.ValidateEmp(employee)
	if err != nil {
		return err
	}
	return s.repository.Save(employee)
}
func (s *EmployeeServiceImpl) GetById(id int) (model.Employee, error) {
	return s.repository.FindById(id)
}
func (s *EmployeeServiceImpl) GetAll() []model.Employee {
	return s.repository.FindAll()
}
func (s *EmployeeServiceImpl) UpdateEmployee(employee model.Employee) error {
	err := utility.ValidateEmp(employee)
	if err != nil {
		return err
	}
	return s.repository.Update(employee)
}
func (s *EmployeeServiceImpl) DeleteEmployee(id int) error {
	return s.repository.Delete(id)
}
