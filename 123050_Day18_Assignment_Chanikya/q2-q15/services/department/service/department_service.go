package service

import (
	"strings"

	"employee-management-microservices/services/department/model"
	"employee-management-microservices/services/department/repository"
	"employee-management-microservices/utility"
)

// Service contains department business logic.
type Service struct {
	repo repository.Repository
}

// New creates the department service.
func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(d *model.Department) error {
	if err := utility.ValidateDepartment(d.Name, d.Description); err != nil {
		return err
	}

	d.Name = strings.TrimSpace(d.Name)
	return s.repo.Create(d)
}

func (s *Service) GetAll() ([]model.Department, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id uint) (*model.Department, error) {
	return s.repo.GetByID(id)
}

func (s *Service) Update(d *model.Department) error {
	if err := utility.ValidateDepartment(d.Name, d.Description); err != nil {
		return err
	}

	d.Name = strings.TrimSpace(d.Name)
	return s.repo.Update(d)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}
