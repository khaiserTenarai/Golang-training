package service

import (
	"employee-platform/model"
	"employee-platform/repository"
	"strings"
)

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}
func (s *DepartmentService) List() ([]model.Department, error)       { return s.repo.List() }
func (s *DepartmentService) ByID(id uint) (*model.Department, error) { return s.repo.ByID(id) }
func (s *DepartmentService) Create(name, description string) (*model.Department, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errDepartmentNameRequired
	}
	d := &model.Department{Name: name, Description: strings.TrimSpace(description)}
	if err := s.repo.Create(d); err != nil {
		return nil, err
	}
	return d, nil
}

var errDepartmentNameRequired = departmentError("department name is required")

type departmentError string

func (e departmentError) Error() string { return string(e) }
