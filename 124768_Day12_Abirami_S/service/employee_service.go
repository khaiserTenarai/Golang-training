package service

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"day12/model"

	"github.com/jackc/pgx/v5"
)

var ErrInvalidInput = errors.New("invalid employee input")

type EmployeeStore interface {
	Create(context.Context, model.EmployeeInput) (model.Employee, error)
	GetByID(context.Context, int64) (model.Employee, error)
	GetAll(context.Context, int, int, string, string, string) ([]model.Employee, int, error)
	Update(context.Context, int64, model.EmployeeInput) (model.Employee, error)
	Delete(context.Context, int64) error
}

type EmployeeService struct{ repo EmployeeStore }

func NewEmployeeService(repo EmployeeStore) *EmployeeService { return &EmployeeService{repo: repo} }

func validate(in model.EmployeeInput) error {
	in.Name, in.Email, in.Department = strings.TrimSpace(in.Name), strings.TrimSpace(in.Email), strings.TrimSpace(in.Department)
	if in.Name == "" || len(in.Name) > 100 || in.Department == "" || len(in.Department) > 100 || in.Email == "" || len(in.Email) > 150 || in.Salary < 0 {
		return ErrInvalidInput
	}
	if _, err := mail.ParseAddress(in.Email); err != nil {
		return ErrInvalidInput
	}
	return nil
}

func (s *EmployeeService) Create(ctx context.Context, in model.EmployeeInput) (model.Employee, error) {
	if err := validate(in); err != nil {
		return model.Employee{}, err
	}
	in.Name, in.Email, in.Department = strings.TrimSpace(in.Name), strings.TrimSpace(in.Email), strings.TrimSpace(in.Department)
	return s.repo.Create(ctx, in)
}
func (s *EmployeeService) GetByID(ctx context.Context, id int64) (model.Employee, error) {
	return s.repo.GetByID(ctx, id)
}
func (s *EmployeeService) GetAll(ctx context.Context, page, limit int, department, sortBy, order string) ([]model.Employee, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if sortBy != "id" && sortBy != "name" && sortBy != "salary" && sortBy != "department" && sortBy != "created_at" {
		sortBy = "id"
	}
	if strings.ToLower(order) != "desc" {
		order = "asc"
	}
	return s.repo.GetAll(ctx, page, limit, strings.TrimSpace(department), sortBy, order)
}
func (s *EmployeeService) Update(ctx context.Context, id int64, in model.EmployeeInput) (model.Employee, error) {
	if err := validate(in); err != nil {
		return model.Employee{}, err
	}
	in.Name, in.Email, in.Department = strings.TrimSpace(in.Name), strings.TrimSpace(in.Email), strings.TrimSpace(in.Department)
	return s.repo.Update(ctx, id, in)
}
func (s *EmployeeService) Delete(ctx context.Context, id int64) error { return s.repo.Delete(ctx, id) }
func IsNotFound(err error) bool                                       { return errors.Is(err, pgx.ErrNoRows) }
