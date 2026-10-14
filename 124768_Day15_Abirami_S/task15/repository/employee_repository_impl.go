package repository

import "day15-secure-gin-api/model"

type EmployeeRepositoryImpl struct {
	employees []model.Employee
	nextID    int
}

func NewEmployeeRepository() *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{[]model.Employee{{1, "Abirami", "abirami@example.com", "IT", 40000}, {2, "Priya", "priya@example.com", "HR", 35000}, {3, "Kavi", "kavi@example.com", "Finance", 45000}}, 4}
}
func (r *EmployeeRepositoryImpl) GetAll() []model.Employee {
	return r.employees
}
func (r *EmployeeRepositoryImpl) GetByID(id int) (model.Employee, bool) {
	for _, e := range r.employees {
		if e.ID == id {
			return e, true
		}
	}
	return model.Employee{}, false
}
func (r *EmployeeRepositoryImpl) Create(e model.Employee) model.Employee {
	e.ID = r.nextID
	r.nextID++
	r.employees = append(r.employees, e)
	return e
}
func (r *EmployeeRepositoryImpl) Update(id int, e model.Employee) (model.Employee, bool) {
	for i := range r.employees {
		if r.employees[i].ID == id {
			e.ID = id
			r.employees[i] = e
			return e, true
		}
	}
	return model.Employee{}, false
}
func (r *EmployeeRepositoryImpl) Delete(id int) bool {
	for i := range r.employees {
		if r.employees[i].ID == id {
			r.employees = append(r.employees[:i], r.employees[i+1:]...)
			return true
		}
	}
	return false
}
