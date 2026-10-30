package service

import (
	"context"
	"employee-service/client"
	"employee-service/dto"
	"employee-service/entity"
	"employee-service/repository"
)

type EmployeeService interface {
	CreateEmployee(ctx context.Context, req dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
	GetEmployeeByID(ctx context.Context, id int64, correlationID string) (*dto.EmployeeResponse, error)
	GetAllEmployees(ctx context.Context) ([]dto.EmployeeResponse, error)
	UpdateEmployee(ctx context.Context, id int64, req dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error)
	DeleteEmployee(ctx context.Context, id int64) error
}

type employeeService struct {
	repo       repository.EmployeeRepository
	deptClient client.DepartmentClient
}

func NewEmployeeService(repo repository.EmployeeRepository, deptClient client.DepartmentClient) EmployeeService {
	return &employeeService{repo: repo, deptClient: deptClient}
}

func (s *employeeService) CreateEmployee(ctx context.Context, req dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	emp := &entity.Employee{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		DepartmentID: req.DepartmentID,
		Status:       entity.StatusPending,
	}

	if err := s.repo.Create(ctx, emp); err != nil {
		return nil, err
	}

	return &dto.EmployeeResponse{
		ID:           emp.ID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		DepartmentID: emp.DepartmentID,
		Status:       string(emp.Status),
		CreatedAt:    emp.CreatedAt,
		UpdatedAt:    emp.UpdatedAt,
	}, nil
}

func (s *employeeService) GetEmployeeByID(ctx context.Context, id int64, correlationID string) (*dto.EmployeeResponse, error) {
	emp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &dto.EmployeeResponse{
		ID:           emp.ID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		DepartmentID: emp.DepartmentID,
		Status:       string(emp.Status),
		CreatedAt:    emp.CreatedAt,
		UpdatedAt:    emp.UpdatedAt,
	}

	if dept, err := s.deptClient.GetDepartmentByID(ctx, emp.DepartmentID, correlationID); err == nil {
		resp.Department = dept
	}

	return resp, nil
}

func (s *employeeService) GetAllEmployees(ctx context.Context) ([]dto.EmployeeResponse, error) {
	emps, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var resps []dto.EmployeeResponse
	for _, emp := range emps {
		resps = append(resps, dto.EmployeeResponse{
			ID:           emp.ID,
			FirstName:    emp.FirstName,
			LastName:     emp.LastName,
			Email:        emp.Email,
			DepartmentID: emp.DepartmentID,
			Status:       string(emp.Status),
			CreatedAt:    emp.CreatedAt,
			UpdatedAt:    emp.UpdatedAt,
		})
	}
	return resps, nil
}

func (s *employeeService) UpdateEmployee(ctx context.Context, id int64, req dto.UpdateEmployeeRequest) (*dto.EmployeeResponse, error) {
	emp, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" { emp.FirstName = req.FirstName }
	if req.LastName != "" { emp.LastName = req.LastName }
	if req.Email != "" { emp.Email = req.Email }
	if req.DepartmentID != 0 { emp.DepartmentID = req.DepartmentID }
	if req.Status != "" { emp.Status = entity.EmployeeStatus(req.Status) }

	if err := s.repo.Update(ctx, emp); err != nil {
		return nil, err
	}

	return &dto.EmployeeResponse{
		ID:           emp.ID,
		FirstName:    emp.FirstName,
		LastName:     emp.LastName,
		Email:        emp.Email,
		DepartmentID: emp.DepartmentID,
		Status:       string(emp.Status),
		CreatedAt:    emp.CreatedAt,
		UpdatedAt:    emp.UpdatedAt,
	}, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
