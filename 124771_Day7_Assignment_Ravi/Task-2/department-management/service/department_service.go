package service

import (
	"context"
	"fmt"
	"strings"

	"department-management/model"
	"department-management/repository"
	"department-management/utility"
)

type DepartmentService struct {
	Repository *repository.DepartmentRepository
}

func NewDepartmentService(
	repository *repository.DepartmentRepository,
) *DepartmentService {

	return &DepartmentService{
		Repository: repository,
	}
}

func (s *DepartmentService) Create(
	ctx context.Context,
	name string,
	description string,
) (*model.Department, error) {

	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if !utility.ValidateDepartmentName(name) {
		return nil, fmt.Errorf(
			"department name must contain at least 2 characters",
		)
	}

	department := &model.Department{
		Name:        name,
		Description: description,
	}

	err := s.Repository.Create(
		ctx,
		department,
	)

	if err != nil {
		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) GetByID(
	ctx context.Context,
	id int,
) (*model.Department, error) {

	if !utility.ValidateID(id) {
		return nil, fmt.Errorf("invalid department ID")
	}

	return s.Repository.GetByID(ctx, id)
}

func (s *DepartmentService) GetAll(
	ctx context.Context,
) ([]model.Department, error) {

	return s.Repository.GetAll(ctx)
}

func (s *DepartmentService) Update(
	ctx context.Context,
	id int,
	name string,
	description string,
) error {

	if !utility.ValidateID(id) {
		return fmt.Errorf("invalid department ID")
	}

	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if !utility.ValidateDepartmentName(name) {
		return fmt.Errorf(
			"department name must contain at least 2 characters",
		)
	}

	department := &model.Department{
		ID:          id,
		Name:        name,
		Description: description,
	}

	return s.Repository.Update(
		ctx,
		department,
	)
}

func (s *DepartmentService) Delete(
	ctx context.Context,
	id int,
) error {

	if !utility.ValidateID(id) {
		return fmt.Errorf("invalid department ID")
	}

	return s.Repository.Delete(ctx, id)
}
