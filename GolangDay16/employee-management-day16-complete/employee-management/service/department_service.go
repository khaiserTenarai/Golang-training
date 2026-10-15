package service

import (
	"employee-management/model"
	"employee-management/repository"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type DepartmentService interface {
	Create(*model.Department) error
	GetAll() ([]model.Department, error)
	GetByID(uint) (*model.Department, error)
}
type departmentService struct{ repo repository.DepartmentRepository }

func NewDepartmentService(r repository.DepartmentRepository) DepartmentService {
	return &departmentService{r}
}
func (s *departmentService) Create(d *model.Department) error {
	d.Name = strings.TrimSpace(d.Name)
	d.Description = strings.TrimSpace(d.Description)
	if d.Name == "" {
		return invalid("department name is required")
	}
	if err := s.repo.Create(d); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return ErrDuplicateDepartment
		}
		return err
	}
	return nil
}
func (s *departmentService) GetAll() ([]model.Department, error) { return s.repo.GetAll() }
func (s *departmentService) GetByID(id uint) (*model.Department, error) {
	return s.repo.GetByID(id)
}
