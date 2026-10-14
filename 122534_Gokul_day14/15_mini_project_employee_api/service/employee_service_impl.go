package service

import (
	"errors"
	"strings"

	"employeegin/dao"
	"employeegin/model"
)

type employeeServiceImpl struct {
	employeeDAO dao.EmployeeDAO
}

func NewEmployeeService(employeeDAO dao.EmployeeDAO) EmployeeService {
	return &employeeServiceImpl{employeeDAO: employeeDAO}
}

func (s *employeeServiceImpl) emailTaken(email string, ownID int) (bool, error) {
	found, err := s.employeeDAO.FindByEmail(email)
	if errors.Is(err, dao.ErrNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return found.ID != ownID, nil
}

func (s *employeeServiceImpl) AddEmployee(req model.EmployeeRequest) (model.Employee, error) {
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	taken, err := s.emailTaken(req.Email, 0)
	if err != nil {
		return model.Employee{}, err
	}
	if taken {
		return model.Employee{}, ErrDuplicateEmail
	}

	return s.employeeDAO.Save(model.Employee{
		Name:       strings.TrimSpace(req.Name),
		Email:      req.Email,
		Department: strings.TrimSpace(req.Department),
		Salary:     req.Salary,
	})
}

func (s *employeeServiceImpl) GetEmployee(id int) (model.Employee, error) {
	return s.employeeDAO.FindByID(id)
}

func (s *employeeServiceImpl) GetAllEmployees(department string) ([]model.Employee, error) {
	return s.employeeDAO.FindAll(department)
}

func (s *employeeServiceImpl) ReplaceEmployee(id int, req model.EmployeeRequest) (model.Employee, error) {
	if _, err := s.employeeDAO.FindByID(id); err != nil {
		return model.Employee{}, err
	}

	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	taken, err := s.emailTaken(req.Email, id)
	if err != nil {
		return model.Employee{}, err
	}
	if taken {
		return model.Employee{}, ErrDuplicateEmail
	}

	updated := model.Employee{
		ID:         id,
		Name:       strings.TrimSpace(req.Name),
		Email:      req.Email,
		Department: strings.TrimSpace(req.Department),
		Salary:     req.Salary,
	}
	if err := s.employeeDAO.Update(updated); err != nil {
		return model.Employee{}, err
	}
	return updated, nil
}

func (s *employeeServiceImpl) PatchEmployee(id int, patch model.EmployeePatch) (model.Employee, error) {
	employee, err := s.employeeDAO.FindByID(id)
	if err != nil {
		return model.Employee{}, err
	}

	if patch.Name != nil {
		employee.Name = strings.TrimSpace(*patch.Name)
	}
	if patch.Department != nil {
		employee.Department = strings.TrimSpace(*patch.Department)
	}
	if patch.Salary != nil {
		employee.Salary = *patch.Salary
	}
	if patch.Email != nil {
		email := strings.ToLower(strings.TrimSpace(*patch.Email))
		taken, err := s.emailTaken(email, id)
		if err != nil {
			return model.Employee{}, err
		}
		if taken {
			return model.Employee{}, ErrDuplicateEmail
		}
		employee.Email = email
	}

	if err := s.employeeDAO.Update(employee); err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (s *employeeServiceImpl) DeleteEmployee(id int) error {
	return s.employeeDAO.Delete(id)
}
