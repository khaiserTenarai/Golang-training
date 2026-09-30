package service

import (
	"department-management/model"
	"department-management/repository"
	"department-management/utility"
)

type DepartmentServiceImpl struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(
	repository repository.DepartmentRepository,
) DepartmentService {

	return &DepartmentServiceImpl{
		repository: repository,
	}
}

func (s *DepartmentServiceImpl) Save(
	department model.Department,
) error {

	err := utility.ValidateDepartment(department)

	if err != nil {
		return err
	}

	return s.repository.Save(department)
}

func (s *DepartmentServiceImpl) FindByID(
	id int,
) (model.Department, error) {

	err := utility.ValidateID(id)

	if err != nil {
		return model.Department{}, err
	}

	return s.repository.FindByID(id)
}

func (s *DepartmentServiceImpl) FindAll() (
	[]model.Department,
	error,
) {

	return s.repository.FindAll()
}

func (s *DepartmentServiceImpl) Update(
	department model.Department,
) error {

	err := utility.ValidateID(department.ID)

	if err != nil {
		return err
	}

	err = utility.ValidateDepartment(department)

	if err != nil {
		return err
	}

	return s.repository.Update(department)
}

func (s *DepartmentServiceImpl) Delete(id int) error {

	err := utility.ValidateID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}
