package service

import (
	"department-management/model"
	"department-management/repository"
	"department-management/utility"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	repository repository.EmployeeRepository,
) EmployeeService {

	return &EmployeeServiceImpl{
		repository: repository,
	}
}

func (s *EmployeeServiceImpl) Save(
	employee model.Employee,
) error {

	err := utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(employee.Email)

	if err != nil {
		return err
	}

	err = utility.ValidateDepartmentID(
		employee.DepartmentID,
	)

	if err != nil {
		return err
	}

	return s.repository.Save(employee)
}

func (s *EmployeeServiceImpl) FindByID(
	id int,
) (model.Employee, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return model.Employee{}, err
	}

	return s.repository.FindByID(id)
}

func (s *EmployeeServiceImpl) FindAll() (
	[]model.Employee,
	error,
) {

	return s.repository.FindAll()
}

func (s *EmployeeServiceImpl) Update(
	employee model.Employee,
) error {

	err := utility.ValidateID(employee.ID)

	if err != nil {
		return err
	}

	err = utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	err = utility.ValidateEmail(employee.Email)

	if err != nil {
		return err
	}

	err = utility.ValidateDepartmentID(
		employee.DepartmentID,
	)

	if err != nil {
		return err
	}

	return s.repository.Update(employee)
}

func (s *EmployeeServiceImpl) Delete(id int) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}
