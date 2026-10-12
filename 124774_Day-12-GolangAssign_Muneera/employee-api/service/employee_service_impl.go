package service

import (
	"context"
	"errors"
	"math"

	"day12/model"
	"day12/repository"
	"day12/validation"
)

type EmployeeServiceImpl struct {
	Repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{Repo: repo}
}

func (s *EmployeeServiceImpl) Create(ctx context.Context, e *model.Employee) error {
	if err := validation.ValidateEmployee(e); err != nil {
		return err
	}
	return s.Repo.Create(ctx, e)
}

func (s *EmployeeServiceImpl) GetAll(
	ctx context.Context, department, name, sortBy, sortOrder string,
	page, pageSize int,
) (*model.EmployeeListResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	if sortOrder != "" && sortOrder != "asc" && sortOrder != "desc" {
		return nil, errors.New("sortOrder must be asc or desc")
	}

	items, total, err := s.Repo.GetAll(
		ctx, department, name, sortBy, sortOrder, page, pageSize,
	)
	if err != nil {
		return nil, err
	}

	return &model.EmployeeListResponse{
		Data:       items,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}, nil
}

func (s *EmployeeServiceImpl) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("invalid employee id")
	}
	return s.Repo.GetByID(ctx, id)
}

func (s *EmployeeServiceImpl) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New("invalid employee id")
	}
	if err := validation.ValidateEmployee(e); err != nil {
		return nil, err
	}
	return s.Repo.Update(ctx, id, e)
}

func (s *EmployeeServiceImpl) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid employee id")
	}
	return s.Repo.Delete(ctx, id)
}
