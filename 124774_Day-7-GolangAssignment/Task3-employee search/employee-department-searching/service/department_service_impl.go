package service

import (
	"errors"

	"employee-management-app/model"
	"employee-management-app/repository"
)

type DepartmentServiceImpl struct {
	repository repository.DepartmentRepository
}

// Constructor
func NewDepartmentService(
	repository repository.DepartmentRepository,
) DepartmentService {

	return &DepartmentServiceImpl{
		repository: repository,
	}
}

// ==================================================
// CREATE
// ==================================================

func (s *DepartmentServiceImpl) AddDepartment(
	department model.Department,
) error {

	if department.Name == "" {
		return errors.New("department name cannot be empty")
	}

	return s.repository.Save(department)
}

// ==================================================
// DELETE
// ==================================================

func (s *DepartmentServiceImpl) DeleteDepartment(
	id int,
) error {

	if id <= 0 {
		return errors.New("department ID must be greater than 0")
	}

	return s.repository.Delete(id)
}

// ==================================================
// UPDATE
// ==================================================

func (s *DepartmentServiceImpl) UpdateDepartment(
	department model.Department,
) error {

	if department.ID <= 0 {
		return errors.New("department ID must be greater than 0")
	}

	if department.Name == "" {
		return errors.New("department name cannot be empty")
	}

	return s.repository.Update(department)
}

// ==================================================
// FIND BY ID
// ==================================================

func (s *DepartmentServiceImpl) FindDepartmentByID(
	id int,
) (model.Department, error) {

	if id <= 0 {
		return model.Department{}, errors.New(
			"department ID must be greater than 0",
		)
	}

	return s.repository.FindByID(id)
}

// ==================================================
// FIND ALL
// ==================================================

func (s *DepartmentServiceImpl) FindAllDepartments() []model.Department {

	return s.repository.FindAll()
}
