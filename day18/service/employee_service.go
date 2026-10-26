package service

import (
	"context"

	"day18/logging"
	"day18/model"
	"day18/repository"
)

type EmployeeService interface {
	OnboardEmployee(ctx context.Context, firstName, lastName, email string, deptID int64) (*model.Employee, error)
	GetEmployee(ctx context.Context, id int64) (*model.Employee, error)
}

type employeeService struct {
	repo   repository.EmployeeRepository
	logger *logging.Logger
}

func NewEmployeeService(repo repository.EmployeeRepository, logger *logging.Logger) EmployeeService {
	return &employeeService{
		repo:   repo,
		logger: logger,
	}
}

func (s *employeeService) OnboardEmployee(ctx context.Context, firstName, lastName, email string, deptID int64) (*model.Employee, error) {
	emp := &model.Employee{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		DepartmentID: deptID,
		Status:       model.StatusPending,
	}

	eventPayload := map[string]interface{}{
		"department_id": emp.DepartmentID,
		"email":         emp.Email,
	}

	err := s.repo.CreateWithOutbox(ctx, emp, "EmployeeOnboardingInitiated", eventPayload)
	if err != nil {
		s.logger.Error("Failed to initiate onboarding saga: %v", err)
		return nil, err
	}

	s.logger.Info("Initiated onboarding process for employee ID: %d", emp.ID)
	return emp, nil
}

func (s *employeeService) GetEmployee(ctx context.Context, id int64) (*model.Employee, error) {
	return s.repo.FindByID(ctx, id)
}
