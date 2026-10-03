package service

import "employee-management/repository"

type SalaryServiceImpl struct {
	repository repository.SalaryRepository
}

func NewSalaryService(repository repository.SalaryRepository) SalaryService {
	return &SalaryServiceImpl{
		repository: repository,
	}
}

func (s *SalaryServiceImpl) UpdateSalary(employeeID int, newSalary float64) error {
	return s.repository.UpdateSalary(employeeID, newSalary)
}
