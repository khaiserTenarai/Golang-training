package service

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"employee-api-day13/model"
	"employee-api-day13/repository"
)

var (
	ErrInvalidInput = errors.New("name, valid email, and department are required")
	ErrFeatureOff   = errors.New("employee creation is disabled by feature flag")
)

type EmployeeService struct {
	repo             repository.EmployeeRepository
	createEmployeeOn bool
}

func NewEmployeeService(repo repository.EmployeeRepository, createEmployeeOn bool) *EmployeeService {
	return &EmployeeService{repo: repo, createEmployeeOn: createEmployeeOn}
}

// Task 11: contexts are passed through the service layer to the repository.
func (s *EmployeeService) Create(ctx context.Context, input model.EmployeeInput) (model.Employee, error) {
	if !s.createEmployeeOn {
		return model.Employee{}, ErrFeatureOff // Task 14: configuration-based feature flag.
	}
	input, err := validateAndNormalize(input)
	if err != nil {
		return model.Employee{}, err
	}
	return s.repo.Create(ctx, input)
}

func (s *EmployeeService) List(ctx context.Context) ([]model.Employee, error) {
	return s.repo.List(ctx)
}

func (s *EmployeeService) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) Update(ctx context.Context, id int64, input model.EmployeeInput) (model.Employee, error) {
	input, err := validateAndNormalize(input)
	if err != nil {
		return model.Employee{}, err
	}
	return s.repo.Update(ctx, id, input)
}

func (s *EmployeeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func validateAndNormalize(input model.EmployeeInput) (model.EmployeeInput, error) {
	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(strings.ToLower(input.Email))
	input.Department = strings.TrimSpace(input.Department)
	if input.Name == "" || input.Department == "" {
		return model.EmployeeInput{}, ErrInvalidInput
	}
	if _, err := mail.ParseAddress(input.Email); err != nil {
		return model.EmployeeInput{}, ErrInvalidInput
	}
	return input, nil
}
