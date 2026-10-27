package repository

import "employee-microservices/department/model"

// DepartmentRepository is an in-memory repository. Tasks 5 and 7.
type DepartmentRepository struct{ items []model.Department }

func NewDepartmentRepository() *DepartmentRepository {
	return &DepartmentRepository{items: []model.Department{
		{ID: 1, Name: "Engineering"},
		{ID: 2, Name: "Human Resources"},
		{ID: 3, Name: "Finance"},
	}}
}
func (r *DepartmentRepository) List() []model.Department { return r.items }
