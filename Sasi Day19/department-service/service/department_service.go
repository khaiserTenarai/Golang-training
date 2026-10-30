package service

import (
	"context"
	"department-service/dto"
	"department-service/entity"
	"department-service/repository"
)

type DepartmentService interface {
	CreateDepartment(ctx context.Context, req dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error)
	GetDepartmentByID(ctx context.Context, id int64) (*dto.DepartmentResponse, error)
	GetAllDepartments(ctx context.Context) ([]dto.DepartmentResponse, error)
}

type departmentService struct {
	repo repository.DepartmentRepository
}

func NewDepartmentService(repo repository.DepartmentRepository) DepartmentService {
	return &departmentService{repo: repo}
}

func (s *departmentService) CreateDepartment(ctx context.Context, req dto.CreateDepartmentRequest) (*dto.DepartmentResponse, error) {
	dept := &entity.Department{
		Name:           req.Name,
		Code:           req.Code,
		HeadcountLimit: req.HeadcountLimit,
	}
	if err := s.repo.Create(ctx, dept); err != nil {
		return nil, err
	}
	return &dto.DepartmentResponse{
		ID:             dept.ID,
		Name:           dept.Name,
		Code:           dept.Code,
		HeadcountLimit: dept.HeadcountLimit,
		CreatedAt:      dept.CreatedAt,
	}, nil
}

func (s *departmentService) GetDepartmentByID(ctx context.Context, id int64) (*dto.DepartmentResponse, error) {
	dept, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.DepartmentResponse{
		ID:               dept.ID,
		Name:             dept.Name,
		Code:             dept.Code,
		HeadcountLimit:   dept.HeadcountLimit,
		CurrentHeadcount: dept.CurrentHeadcount,
		CreatedAt:        dept.CreatedAt,
	}, nil
}

func (s *departmentService) GetAllDepartments(ctx context.Context) ([]dto.DepartmentResponse, error) {
	depts, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var resps []dto.DepartmentResponse
	for _, dept := range depts {
		resps = append(resps, dto.DepartmentResponse{
			ID:               dept.ID,
			Name:             dept.Name,
			Code:             dept.Code,
			HeadcountLimit:   dept.HeadcountLimit,
			CurrentHeadcount: dept.CurrentHeadcount,
			CreatedAt:        dept.CreatedAt,
		})
	}
	return resps, nil
}
