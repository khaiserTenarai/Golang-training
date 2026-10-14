package service

import (
	"context"
	"errors"
	"strings"

	"day-14-gin-employee-api/model"
	"day-14-gin-employee-api/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(ctx context.Context, e *model.Employee) error {
	if err := validate(e); err != nil {
		return err
	}
	normalize(e)
	return s.repo.Create(ctx, e)
}

func (s *EmployeeService) GetAll(ctx context.Context) ([]model.Employee, error) {
	return s.repo.GetAll(ctx)
}

func (s *EmployeeService) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *EmployeeService) Update(ctx context.Context, id int64, e *model.Employee) (model.Employee, error) {
	if err := validate(e); err != nil {
		return model.Employee{}, err
	}
	normalize(e)
	return s.repo.Update(ctx, id, e)
}

func (s *EmployeeService) Patch(ctx context.Context, id int64, fields map[string]any) (model.Employee, error) {
	for key, value := range fields {
		switch key {
		case "name", "email", "department":
			str, ok := value.(string)
			if !ok || strings.TrimSpace(str) == "" {
				return model.Employee{}, errors.New(key + " must be a non-empty string")
			}
			fields[key] = strings.TrimSpace(str)
		case "salary":
			switch value.(type) {
			case float64:
				if value.(float64) < 0 {
					return model.Employee{}, errors.New("salary cannot be negative")
				}
			default:
				return model.Employee{}, errors.New("salary must be a number")
			}
		default:
			return model.Employee{}, errors.New("unsupported field: " + key)
		}
	}
	if len(fields) == 0 {
		return model.Employee{}, errors.New("provide at least one field")
	}
	return s.repo.Patch(ctx, id, fields)
}

func (s *EmployeeService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func normalize(e *model.Employee) {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.TrimSpace(e.Email)
	e.Department = strings.TrimSpace(e.Department)
}

func validate(e *model.Employee) error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(strings.TrimSpace(e.Email), "@") {
		return errors.New("a valid email is required")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("department is required")
	}
	if e.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}
