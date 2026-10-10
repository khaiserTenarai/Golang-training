package service

import (
	"student-management/model"
	"student-management/repository"
)

// StudentService handles business logic.
type StudentService struct {
	repo *repository.StudentRepository
}

// NewStudentService creates the service.
func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

// GetAll returns all students.
func (s *StudentService) GetAll() []model.Student {
	return s.repo.GetAll()
}

// Create adds a student.
func (s *StudentService) Create(student model.Student) model.Student {
	return s.repo.Create(student)
}
