package service

import (
	"context"
	"errors"
	"strings"

	"student-management/model"
	"student-management/repository"
	"student-management/utility"
)

// StudentService contains business logic.
type StudentService struct {
	repo repository.StudentRepository
}

// NewStudentService creates the service.
func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

// GetStudents filters and paginates students.
func (s *StudentService) GetStudents(ctx context.Context, name, grade string, page, limit int) ([]model.Student, int64, error) {
	if page < 1 {
		return nil, 0, errors.New("page must be at least 1")
	}
	if limit < 1 || limit > 100 {
		return nil, 0, errors.New("limit must be between 1 and 100")
	}

	name = strings.TrimSpace(name)
	grade = strings.TrimSpace(grade)
	offset := (page - 1) * limit

	students, err := s.repo.GetFiltered(ctx, name, grade, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.CountFiltered(ctx, name, grade)
	if err != nil {
		return nil, 0, err
	}
	return students, total, nil
}

// CreateStudent validates and saves a student.
func (s *StudentService) CreateStudent(ctx context.Context, student model.Student) (model.Student, error) {
	student.Name = strings.TrimSpace(student.Name)
	student.Grade = strings.TrimSpace(student.Grade)
	if err := utility.ValidateStudent(student); err != nil {
		return model.Student{}, err
	}
	return s.repo.Create(ctx, student)
}
