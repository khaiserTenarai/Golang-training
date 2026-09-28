package main

type EmployeeService interface {
	AddEmployee(e Employee) (Employee, error)
	GetEmployee(id int) (Employee, error)
	ListEmployees() []Employee
	RaiseSalary(id int, amount float64) (Employee, error)
	DeactivateEmployee(id int) error
	RemoveEmployee(id int) error
}

type employeeService struct {
	repo EmployeeRepository
}

func NewEmployeeService(repo EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (s *employeeService) AddEmployee(e Employee) (Employee, error) {
	e.Active = true
	return s.repo.Create(e)
}

func (s *employeeService) GetEmployee(id int) (Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) ListEmployees() []Employee {
	return s.repo.GetAll()
}

func (s *employeeService) RaiseSalary(id int, amount float64) (Employee, error) {
	e, err := s.repo.GetByID(id)
	if err != nil {
		return Employee{}, err
	}
	e.GiveRaise(amount)
	if err := s.repo.Update(e); err != nil {
		return Employee{}, err
	}
	return e, nil
}

func (s *employeeService) DeactivateEmployee(id int) error {
	e, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	e.Deactivate()
	return s.repo.Update(e)
}

func (s *employeeService) RemoveEmployee(id int) error {
	return s.repo.Delete(id)
}
