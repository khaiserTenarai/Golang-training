package repository

import "student-management/model"

// MemoryRepository stores students in memory.
type MemoryRepository struct {
	students []model.Student
	nextID   int
}

// NewMemoryRepository creates the repository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		students: []model.Student{},
		nextID:   1,
	}
}

// GetAll returns all students.
func (r *MemoryRepository) GetAll() []model.Student {
	return r.students
}

// Create saves a student.
func (r *MemoryRepository) Create(student model.Student) model.Student {
	student.ID = r.nextID
	r.nextID++

	r.students = append(r.students, student)

	return student
}
