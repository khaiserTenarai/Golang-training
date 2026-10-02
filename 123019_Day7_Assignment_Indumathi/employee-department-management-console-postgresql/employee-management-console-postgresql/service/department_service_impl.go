package service

import (
	"errors"
	"strings"

	"example.com/employee-management/models"
	"example.com/employee-management/repository"
)

type departmentServiceImpl struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentServiceImpl{
		repository: repo,
	}
}

func (s *departmentServiceImpl) validate(dept models.Department) error {
	if strings.TrimSpace(dept.Name) == "" {
		return errors.New("department name is required")
	}
	if strings.TrimSpace(dept.Code) == "" {
		return errors.New("department code is required")
	}
	return nil
}

func (s *departmentServiceImpl) AddDepartment(dept models.Department) error {
	if err := s.validate(dept); err != nil {
		return err
	}
	return s.repository.Create(&dept)
}

func (s *departmentServiceImpl) GetDepartment(id int) (models.Department, error) {
	if id <= 0 {
		return models.Department{}, errors.New("invalid department id")
	}
	dept, err := s.repository.FindByID(id)
	if err != nil {
		return models.Department{}, err
	}
	return *dept, nil
}

func (s *departmentServiceImpl) GetAllDepartments() ([]models.Department, error) {
	return s.repository.FindAll()
}

func (s *departmentServiceImpl) UpdateDepartment(dept models.Department) error {
	if dept.ID <= 0 {
		return errors.New("invalid department id")
	}
	if err := s.validate(dept); err != nil {
		return err
	}
	return s.repository.Update(&dept)
}

func (s *departmentServiceImpl) DeleteDepartment(id int) error {
	if id <= 0 {
		return errors.New("invalid department id")
	}
	return s.repository.Delete(id)
}