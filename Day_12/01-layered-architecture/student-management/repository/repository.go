package repository

import "student-management/model"

// StudentRepository stores student data.
type StudentRepository struct {
	students []model.Student
	nextID   int
}

// NewStudentRepository creates the repository.
func NewStudentRepository() *StudentRepository {
	return &StudentRepository{
		students: []model.Student{},
		nextID:   1,
	}
}

// GetAll returns all students.
func (r *StudentRepository) GetAll() []model.Student {
	return r.students
}

// Create saves a student.
func (r *StudentRepository) Create(student model.Student) model.Student {
	student.ID = r.nextID
	r.nextID++

	r.students = append(r.students, student)

	return student
}
