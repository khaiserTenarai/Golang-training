package service

import (
	"fmt"
	"employee_management/models"
	"employee_management/dao"
	"employee_management/utils"

)

type EmployeeServiceImpl struct{
	dao dao.EmployeeDao
}

func NewEmployeeService(dao dao.EmployeeDao) EmployeeService{
	return &EmployeeServiceImpl{dao: dao}
}

func (s *EmployeeServiceImpl) ValidateEmployee (emp models.Employee) error{
	if err := utils.ValidateID(emp.ID); err != nil{
		return err
	}

	if err := utils.ValidateName(emp.Name); err != nil{
		return err
	}

	if err := utils.ValidateAge(emp.Age); err != nil{
		return err
	}
	
	return nil
}

func (s *EmployeeServiceImpl) AddEmployee(employee models.Employee) error{
	fmt.Println("Hello from service - Add Employee")

	if err := s.ValidateEmployee(employee); err != nil{
		return fmt.Errorf("Validation Error: %w", err)
	}

	return s.dao.Save(employee)
}

func (s *EmployeeServiceImpl) GetEmployee(id int) (models.Employee, error){
	fmt.Println("Hello from Service - Get Employee")

	if err := utils.ValidateID(id); err != nil{
		return models.Employee{}, err
	}

	return s.dao.FindById(id)

}

func (s *EmployeeServiceImpl) GetAllEmployee() ([]models.Employee, error){
	fmt.Println("Hello from Service: Get all emp")

	return s.dao.FindAll()
}

func (s *EmployeeServiceImpl) UpdateEmployee(employee models.Employee) error{
	fmt.Println("Hello from Service: Update Employee")

	if err := s.ValidateEmployee(employee); err != nil{
		return fmt.Errorf("Validation Error: %w", err)
	}

	return s.dao.Update(employee)

}

func (s *EmployeeServiceImpl) DeleteEmployee(id int) error{
	fmt.Println("Hello from service: Delete employee")

	if err := utils.ValidateID(id); err != nil{
		return err
	}

	return s.dao.Delete(id)
}