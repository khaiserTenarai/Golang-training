package controller

import (
	"employee-management/model"
	"employee-management/service"
)

type DepartmentControllerImpl struct {
	service service.DepartmentService
}

func NewDepartmentController(service service.DepartmentService) DepartmentController {
	return &DepartmentControllerImpl{
		service: service,
	}
}
func (c *DepartmentControllerImpl) CreateDepartment(department model.Department) error {
	return c.service.CreateDepartment(department)
}
func (c *DepartmentControllerImpl) GetDepartment(id int) (*model.Department, error) {
	return c.service.GetDepartment(id)
}
func (c *DepartmentControllerImpl) GetAllDepartments() ([]model.Department, error) {
	return c.service.GetAllDepartments()
}
func (c *DepartmentControllerImpl) UpdateDepartment(department model.Department) error {
	return c.service.UpdateDepartment(department)
}
func (c *DepartmentControllerImpl) DeleteDepartment(id int) error {
	return c.service.DeleteDepartment(id)
}
