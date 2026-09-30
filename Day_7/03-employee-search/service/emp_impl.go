package service
import (
	"employee-management/model"
	"employee-management/repository"
	"employee-management/utility"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	repository repository.EmployeeRepository,
) EmployeeService {

	return &EmployeeServiceImpl{
		repository: repository,
	}
}

func (s *EmployeeServiceImpl) Save(
	employee model.Employee,
) error {

	err := utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	return s.repository.Save(employee)
}

func (s *EmployeeServiceImpl) FindByID(
	id int,
) (model.Employee, error) {

	return s.repository.FindByID(id)
}

func (s *EmployeeServiceImpl) FindAll() (
	[]model.Employee,
	error,
) {

	return s.repository.FindAll()
}

func (s *EmployeeServiceImpl) Update(
	employee model.Employee,
) error {

	err := utility.ValidateEmployee(employee)

	if err != nil {
		return err
	}

	return s.repository.Update(employee)
}

func (s *EmployeeServiceImpl) Delete(
	id int,
) error {

	if id <= 0 {
		return utility.InvalidEmployeeID
	}

	return s.repository.Delete(id)
}

func (s *EmployeeServiceImpl) Search(
	search model.EmployeeSearch,
) ([]model.Employee, int, error) {

	if search.Page <= 0 {
		search.Page = 1
	}

	if search.Size <= 0 {
		search.Size = 5
	}

	if search.SortBy == "" {
		search.SortBy = "id"
	}

	if search.SortOrder == "" {
		search.SortOrder = "asc"
	}

	return s.repository.Search(search)
}
