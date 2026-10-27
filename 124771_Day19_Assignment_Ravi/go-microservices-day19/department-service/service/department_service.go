package service

import (
	"context"
	"day19/department-service/dto"
	"day19/department-service/model"
	"day19/department-service/repository"
	"errors"
	"gorm.io/gorm"
)

// DepartmentService contains business rules and maps DTOs to/from entities.
type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}
func (s *DepartmentService) Create(ctx context.Context, in dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error) {
	d := &model.Department{Name: in.Name, Description: in.Description}
	if err := s.repo.Create(ctx, d); err != nil {
		return nil, err
	}
	return toResponse(d), nil
}
func (s *DepartmentService) List(ctx context.Context) ([]dto.DepartmentResponse, error) {
	rows, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]dto.DepartmentResponse, 0, len(rows))
	for i := range rows {
		out = append(out, *toResponse(&rows[i]))
	}
	return out, nil
}
func (s *DepartmentService) Get(ctx context.Context, id uint) (*dto.DepartmentResponse, error) {
	d, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return toResponse(d), nil
}
func (s *DepartmentService) Update(ctx context.Context, id uint, in dto.UpdateDepartmentRequest) (*dto.DepartmentResponse, error) {
	d, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	d.Name = in.Name
	d.Description = in.Description
	if err := s.repo.Update(ctx, d); err != nil {
		return nil, err
	}
	return toResponse(d), nil
}
func (s *DepartmentService) Delete(ctx context.Context, id uint) error {
	if _, err := s.repo.Get(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return s.repo.Delete(ctx, id)
}
func toResponse(d *model.Department) *dto.DepartmentResponse {
	return &dto.DepartmentResponse{ID: d.ID, Name: d.Name, Description: d.Description}
}
