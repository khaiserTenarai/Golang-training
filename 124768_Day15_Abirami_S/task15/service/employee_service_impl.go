package service

import (
	"day15-secure-gin-api/model"
	"day15-secure-gin-api/repository"
)

type EmployeeServiceImpl struct{ repo repository.EmployeeRepository }

func NewEmployeeService(r repository.EmployeeRepository) *EmployeeServiceImpl {
	return &EmployeeServiceImpl{r}
}
func (s *EmployeeServiceImpl) GetAll(p, l int) []model.Employee {
	a := s.repo.GetAll()
	if p < 1 {
		p = 1
	}
	if l < 1 {
		l = 2
	}
	start := (p - 1) * l
	if start >= len(a) {
		return []model.Employee{}
	}
	end := start + l
	if end > len(a) {
		end = len(a)
	}
	return a[start:end]
}
func (s *EmployeeServiceImpl) GetByID(id int) (model.Employee, bool) {
	return s.repo.GetByID(id)
}
func (s *EmployeeServiceImpl) Create(e model.Employee) model.Employee {
	return s.repo.Create(e)
}
func (s *EmployeeServiceImpl) Update(id int, e model.Employee) (model.Employee, bool) {
	return s.repo.Update(id, e)
}
func (s *EmployeeServiceImpl) Delete(id int) bool {
	return s.repo.Delete(id)
}
