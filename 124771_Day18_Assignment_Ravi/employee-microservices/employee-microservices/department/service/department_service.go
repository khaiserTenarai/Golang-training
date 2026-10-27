package service

import (
	"employee-microservices/department/model"
	"employee-microservices/department/repository"
)

// DepartmentService contains Department business logic. Task 5.
type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(r *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: r}
}
func (s *DepartmentService) List() []model.Department { return s.repo.List() }
