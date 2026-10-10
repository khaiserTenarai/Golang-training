package service

import (
	"context"
	"errors"
	"strings"

	"student-management/model"
	"student-management/repository"
)

type StudentService interface {
	CreateStudent(context.Context, model.Student) error
	GetStudents(context.Context, string, string, string, string, int, int) ([]model.Student, int, error)
}

type studentService struct{ repo repository.StudentRepository }

func NewStudentService(repo repository.StudentRepository) StudentService {
	return &studentService{repo: repo}
}

func (s *studentService) CreateStudent(ctx context.Context, student model.Student) error {
	student.Name = strings.TrimSpace(student.Name)
	student.Grade = strings.TrimSpace(student.Grade)
	if student.Name == "" {
		return errors.New("student name is required")
	}
	if student.Grade == "" {
		return errors.New("student grade is required")
	}
	return s.repo.Create(ctx, student)
}

func (s *studentService) GetStudents(ctx context.Context, name, grade, sortBy, order string, page, limit int) ([]model.Student, int, error) {
	sortBy = strings.ToLower(sortBy)
	order = strings.ToLower(order)
	if sortBy == "" {
		sortBy = "id"
	}
	if order == "" {
		order = "asc"
	}
	switch sortBy {
	case "id", "name", "grade":
	default:
		return nil, 0, errors.New("sortBy must be id, name, or grade")
	}
	if order != "asc" && order != "desc" {
		return nil, 0, errors.New("order must be asc or desc")
	}
	if page < 1 {
		return nil, 0, errors.New("page must be greater than 0")
	}
	if limit < 1 || limit > 100 {
		return nil, 0, errors.New("limit must be between 1 and 100")
	}
	students, err := s.repo.GetFilteredSorted(ctx, name, grade, sortBy, order, limit, (page-1)*limit)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.CountFiltered(ctx, name, grade)
	if err != nil {
		return nil, 0, err
	}
	return students, total, nil
}
