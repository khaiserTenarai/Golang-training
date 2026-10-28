package service

import (
	"employee-department-microservices/department-service/model"
	"employee-department-microservices/department-service/repository"
)

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(r *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: r}
}
func (s *DepartmentService) List() []model.Department                   { return s.repo.List() }
func (s *DepartmentService) Get(id int) (model.Department, error)       { return s.repo.Get(id) }
func (s *DepartmentService) Create(v model.Department) model.Department { return s.repo.Create(v) }
func (s *DepartmentService) Update(id int, v model.Department) (model.Department, error) {
	return s.repo.Update(id, v)
}
func (s *DepartmentService) Delete(id int) error { return s.repo.Delete(id) }
