package service

import (
	"context"
	"errors"
	"example.com/q12-repository-service/model"
	"example.com/q12-repository-service/repository"
)

type EmployeeServiceImpl struct{ repo repository.EmployeeRepository }

func NewEmployeeService(r repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repo: r}
}
func (s *EmployeeServiceImpl) Add(c context.Context, e model.Employee) error {
	if e.Name == "" || e.Email == "" || e.Department == "" {
		return errors.New("all fields are required")
	}
	return s.repo.Create(c, e)
}
func (s *EmployeeServiceImpl) List(c context.Context) ([]model.Employee, error) {
	return s.repo.FindAll(c)
}
