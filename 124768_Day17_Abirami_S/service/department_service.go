package service

import (
	"day17-employee-platform/model"
	"day17-employee-platform/repository"
)

type DepartmentService interface {
	GetByID(id uint) (*model.Department, error)
}

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) GetByID(id uint) (*model.Department, error) {
	return s.repo.GetByID(id)
}
