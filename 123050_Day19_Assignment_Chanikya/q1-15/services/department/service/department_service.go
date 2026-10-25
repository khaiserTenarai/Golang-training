package service

import (
	"errors"

	"employee-management-microservices/services/department/entity"
	"employee-management-microservices/services/department/repository"

	"gorm.io/gorm"
)

type DepartmentService struct {
	Repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{Repo: repo}
}

func (s *DepartmentService) Create(name string) (*entity.Department, error) {
	department := &entity.Department{Name: name}
	if err := s.Repo.Create(department); err != nil {
		return nil, err
	}
	return department, nil
}

func (s *DepartmentService) GetAll() ([]entity.Department, error) {
	return s.Repo.GetAll()
}

func (s *DepartmentService) GetByID(id uint) (*entity.Department, error) {
	return s.Repo.GetByID(id)
}

func (s *DepartmentService) Update(id uint, name string) (*entity.Department, error) {
	department, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	department.Name = name
	if err := s.Repo.Update(department); err != nil {
		return nil, err
	}
	return department, nil
}

func (s *DepartmentService) Delete(id uint) error {
	err := s.Repo.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return err
}
