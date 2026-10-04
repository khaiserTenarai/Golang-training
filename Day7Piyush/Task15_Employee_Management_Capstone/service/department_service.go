package service

import (
	"errors"
	"task15_employee_management_capstone/models"
	"task15_employee_management_capstone/repository"
)

type DepartmentService interface {
	CreateDepartment(dept models.Department) (models.Department, error)
	GetAllDepartments() ([]models.Department, error)
	GetDepartmentByID(id int) (models.Department, error)
	UpdateDepartment(dept models.Department) (models.Department, error)
	DeleteDepartment(id int) error
}

type DepartmentServiceImpl struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &DepartmentServiceImpl{repo: repo}
}

func (s *DepartmentServiceImpl) CreateDepartment(dept models.Department) (models.Department, error) {
	if dept.Name == "" {
		return models.Department{}, errors.New("department name is required")
	}
	return s.repo.Create(dept)
}

func (s *DepartmentServiceImpl) GetAllDepartments() ([]models.Department, error) {
	return s.repo.GetAll()
}

func (s *DepartmentServiceImpl) GetDepartmentByID(id int) (models.Department, error) {
	if id <= 0 {
		return models.Department{}, errors.New("invalid department ID")
	}
	return s.repo.GetByID(id)
}

func (s *DepartmentServiceImpl) UpdateDepartment(dept models.Department) (models.Department, error) {
	if dept.ID <= 0 {
		return models.Department{}, errors.New("invalid department ID")
	}
	if dept.Name == "" {
		return models.Department{}, errors.New("department name is required")
	}
	return s.repo.Update(dept)
}

func (s *DepartmentServiceImpl) DeleteDepartment(id int) error {
	if id <= 0 {
		return errors.New("invalid department ID")
	}
	return s.repo.Delete(id)
}
