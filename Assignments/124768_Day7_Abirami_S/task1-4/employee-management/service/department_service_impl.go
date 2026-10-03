package service

import (
	"employee-management/model"
	"employee-management/repository"
)

type DepartmentServiceImpl struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(repository repository.DepartmentRepository) DepartmentService {
	return &DepartmentServiceImpl{
		repository: repository,
	}
}
func (s *DepartmentServiceImpl) CreateDepartment(department model.Department) error {
	return s.repository.CreateDepartment(department)
}
func (s *DepartmentServiceImpl) GetDepartment(id int) (*model.Department, error) {
	return s.repository.GetDepartment(id)
}
func (s *DepartmentServiceImpl) GetAllDepartments() ([]model.Department, error) {

	return s.repository.GetAllDepartments()
}
func (s *DepartmentServiceImpl) UpdateDepartment(department model.Department) error {
	return s.repository.UpdateDepartment(department)
}
func (s *DepartmentServiceImpl) DeleteDepartment(id int) error {
	return s.repository.DeleteDepartment(id)
}
