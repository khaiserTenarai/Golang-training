package controller

import "employee-management/service"

type SalaryControllerImpl struct {
	service service.SalaryService
}

func NewSalaryController(service service.SalaryService) SalaryController {
	return &SalaryControllerImpl{
		service: service,
	}
}

func (c *SalaryControllerImpl) UpdateSalary(employeeID int, newSalary float64) error {
	return c.service.UpdateSalary(employeeID, newSalary)
}
