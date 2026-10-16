package service

import (
	"employee-management/model"
	"employee-management/repository"
	"errors"
	"gorm.io/gorm"
	"strings"
	"unicode/utf8"
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
	if utf8.RuneCountInString(d.Name) > 100 {
		return invalid("department name must be at most 100 characters")
	}
	if utf8.RuneCountInString(d.Description) > 255 {
		return invalid("description must be at most 255 characters")
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
