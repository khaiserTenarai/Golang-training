package service

import (
	"context"
	"day19/employee-service/client"
	"day19/employee-service/dto"
	"day19/employee-service/model"
	"day19/employee-service/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// EmployeeService combines local persistence with calls to Department Service.
type EmployeeService struct {
	repo        *repository.EmployeeRepository
	departments *client.DepartmentClient
}

func NewEmployeeService(repo *repository.EmployeeRepository, departments *client.DepartmentClient) *EmployeeService {
	return &EmployeeService{repo: repo, departments: departments}
}
func (s *EmployeeService) Create(ctx context.Context, in dto.CreateEmployeeRequest, cid string) (*dto.EmployeeResponse, error) {
	// Validate the foreign service reference over HTTP before storing the employee.
	if _, err := s.departments.Get(ctx, in.DepartmentID, cid); err != nil {
		return nil, fmt.Errorf("validate department: %w", err)
	}
	e := &model.Employee{Name: in.Name, Email: in.Email, DepartmentID: in.DepartmentID}
	if err := s.repo.Create(ctx, e); err != nil {
		return nil, err
	}
	return s.toResponse(ctx, e, cid)
}
func (s *EmployeeService) List(ctx context.Context, cid string) ([]dto.EmployeeResponse, error) {
	rows, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]dto.EmployeeResponse, 0, len(rows))
	for i := range rows {
		item, err := s.toResponse(ctx, &rows[i], cid)
		if err != nil {
			return nil, err
		}
		out = append(out, *item)
	}
	return out, nil
}
func (s *EmployeeService) Get(ctx context.Context, id uint, cid string) (*dto.EmployeeResponse, error) {
	e, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(ctx, e, cid)
}
func (s *EmployeeService) Update(ctx context.Context, id uint, in dto.UpdateEmployeeRequest, cid string) (*dto.EmployeeResponse, error) {
	if _, err := s.departments.Get(ctx, in.DepartmentID, cid); err != nil {
		return nil, fmt.Errorf("validate department: %w", err)
	}
	e, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	e.Name = in.Name
	e.Email = in.Email
	e.DepartmentID = in.DepartmentID
	if err := s.repo.Update(ctx, e); err != nil {
		return nil, err
	}
	return s.toResponse(ctx, e, cid)
}
func (s *EmployeeService) Delete(ctx context.Context, id uint) error {
	if _, err := s.repo.Get(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return s.repo.Delete(ctx, id)
}
func (s *EmployeeService) toResponse(ctx context.Context, e *model.Employee, cid string) (*dto.EmployeeResponse, error) {
	// Interconnection: each response obtains department data from the other service.
	department, err := s.departments.Get(ctx, e.DepartmentID, cid)
	if err != nil {
		return nil, fmt.Errorf("load department details: %w", err)
	}
	return &dto.EmployeeResponse{ID: e.ID, Name: e.Name, Email: e.Email, DepartmentID: e.DepartmentID, Department: department}, nil
}
