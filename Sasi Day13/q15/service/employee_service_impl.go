package service

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"employeeapi/config"
	"employeeapi/dao"
	"employeeapi/model"
)

type employeeServiceImpl struct {
	employeeDAO dao.EmployeeDAO
	features    config.Features
	log         *slog.Logger
}

func NewEmployeeService(employeeDAO dao.EmployeeDAO, features config.Features, log *slog.Logger) EmployeeService {
	return &employeeServiceImpl{employeeDAO: employeeDAO, features: features, log: log}
}

func validate(employee model.Employee) error {
	if strings.TrimSpace(employee.Name) == "" {
		return fmt.Errorf("%w: name is required", ErrValidation)
	}
	if strings.TrimSpace(employee.Department) == "" {
		return fmt.Errorf("%w: department is required", ErrValidation)
	}
	if employee.Salary < 0 {
		return fmt.Errorf("%w: salary cannot be negative", ErrValidation)
	}
	return nil
}

func (s *employeeServiceImpl) AddEmployee(ctx context.Context, employee model.Employee) (model.Employee, error) {
	if err := validate(employee); err != nil {
		return model.Employee{}, err
	}

	saved, err := s.employeeDAO.Save(ctx, employee)
	if err != nil {
		return model.Employee{}, err
	}

	s.log.Info("employee created", "id", saved.ID, "name", saved.Name)
	return saved, nil
}

func (s *employeeServiceImpl) GetEmployee(ctx context.Context, id int) (model.Employee, error) {
	return s.employeeDAO.FindByID(ctx, id)
}

func (s *employeeServiceImpl) GetAllEmployees(ctx context.Context) ([]model.Employee, error) {
	employees, err := s.employeeDAO.FindAll(ctx)
	if err == nil {
		s.log.Debug("listing employees", "count", len(employees))
	}
	return employees, err
}

func (s *employeeServiceImpl) DeleteEmployee(ctx context.Context, id int) error {
	if !s.features.DeleteEmployee {
		return fmt.Errorf("delete employee: %w", ErrFeatureDisabled)
	}

	if err := s.employeeDAO.Delete(ctx, id); err != nil {
		return err
	}

	s.log.Info("employee deleted", "id", id)
	return nil
}

func (s *employeeServiceImpl) SalaryReport(ctx context.Context) (model.SalaryReport, error) {
	if !s.features.SalaryReport {
		return model.SalaryReport{}, fmt.Errorf("salary report: %w", ErrFeatureDisabled)
	}

	employees, err := s.employeeDAO.FindAll(ctx)
	if err != nil {
		return model.SalaryReport{}, err
	}

	report := model.SalaryReport{Employees: len(employees)}
	for _, e := range employees {
		report.TotalSalary += e.Salary
	}
	if report.Employees > 0 {
		report.AverageSalary = report.TotalSalary / float64(report.Employees)
	}
	return report, nil
}
