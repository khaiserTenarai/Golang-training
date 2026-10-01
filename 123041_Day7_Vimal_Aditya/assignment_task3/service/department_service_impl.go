package service

import (
	"errors"
	"strings"

	"example.com/employee-management/model"
	"example.com/employee-management/repository"
)

type departmentServiceImpl struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(
	repository repository.DepartmentRepository,
) DepartmentService {
	return &departmentServiceImpl{
		repository: repository,
	}
}

func (s *departmentServiceImpl) validate(department model.Department) error {
	if strings.TrimSpace(department.Name) == "" {
		return errors.New("department name is required")
	}
	if strings.TrimSpace(department.Code) == "" {
		return errors.New("department code is required")
	}
	return nil
}

func (s *departmentServiceImpl) AddDepartment(department *model.Department) error {
	if err := s.validate(*department); err != nil {
		return err
	}
	return s.repository.Save(department)
}

func (s *departmentServiceImpl) GetDepartment(id int64) (model.Department, error) {
	if id <= 0 {
		return model.Department{}, errors.New("invalid department id")
	}
	return s.repository.FindByID(id)
}

func (s *departmentServiceImpl) GetAllDepartments() (
	[]model.Department,
	error,
) {
	return s.repository.FindAll()
}

func (s *departmentServiceImpl) UpdateDepartment(department model.Department) error {
	if department.ID <= 0 {
		return errors.New("invalid department id")
	}
	if err := s.validate(department); err != nil {
		return err
	}
	return s.repository.Update(department)
}

func (s *departmentServiceImpl) DeleteDepartment(id int64) error {
	if id <= 0 {
		return errors.New("invalid department id")
	}
	return s.repository.Delete(id)
}