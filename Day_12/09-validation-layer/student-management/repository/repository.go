package repository

import "student-management/model"

// StudentRepository defines data operations.
type StudentRepository interface {
	GetAll() []model.Student
	Create(student model.Student) model.Student
}
