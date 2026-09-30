package service
import (
	"salary-management/model"
	"salary-management/repository"
	"salary-management/utility"
)

type SalaryService interface {

	UpdateSalary(
		employeeID int,
		newSalary float64,
	) error

	FindHistory(
		employeeID int,
	) ([]model.SalaryHistory, error)
}

type SalaryServiceImpl struct {
	repository repository.SalaryRepository
}

func NewSalaryService(
	repository repository.SalaryRepository,
) SalaryService {

	return &SalaryServiceImpl{
		repository: repository,
	}
}

func (s *SalaryServiceImpl) UpdateSalary(
	employeeID int,
	newSalary float64,
) error {

	err := utility.ValidateEmployeeID(
		employeeID,
	)

	if err != nil {
		return err
	}

	err = utility.ValidateSalary(
		newSalary,
	)

	if err != nil {
		return err
	}

	return s.repository.UpdateSalary(
		employeeID,
		newSalary,
	)
}

func (s *SalaryServiceImpl) FindHistory(
	employeeID int,
) ([]model.SalaryHistory, error) {

	err := utility.ValidateEmployeeID(
		employeeID,
	)

	if err != nil {
		return nil, err
	}

	return s.repository.FindHistory(
		employeeID,
	)
}
