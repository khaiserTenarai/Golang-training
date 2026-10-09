package service

import (
	"context"
	"math"

	"employee-management/repository"
	"employee-management/utility"
	"employee-management/view"
)

type EmployeeServiceImpl struct {
	repository repository.EmployeeRepository
}

func NewEmployeeService(
	repository repository.EmployeeRepository,
) *EmployeeServiceImpl {

	return &EmployeeServiceImpl{
		repository: repository,
	}
}

func (s *EmployeeServiceImpl) SearchEmployees(
	ctx context.Context,
	req view.EmployeeSearchRequest,
) (*view.EmployeeSearchResponse, error) {

	// --------------------------------------------
	// Validate Pagination
	// --------------------------------------------

	if err := utility.ValidatePagination(
		req.Page,
		req.PageSize,
	); err != nil {

		return nil, err
	}

	// --------------------------------------------
	// Validate Sorting
	// --------------------------------------------

	if err := utility.ValidateSort(
		req.SortBy,
		req.SortOrder,
	); err != nil {

		return nil, err
	}

	// --------------------------------------------
	// Validate Salary
	// --------------------------------------------

	if err := utility.ValidateSalary(
		req.SalaryMin,
		req.SalaryMax,
	); err != nil {

		return nil, err
	}

	// --------------------------------------------
	// Repository
	// --------------------------------------------

	employees, total, err :=
		s.repository.Search(
			ctx,
			req,
		)

	if err != nil {
		return nil, err
	}

	// --------------------------------------------
	// Total Pages
	// --------------------------------------------

	totalPages := int(
		math.Ceil(
			float64(total) /
				float64(req.PageSize),
		),
	)

	// --------------------------------------------
	// Response
	// --------------------------------------------

	return &view.EmployeeSearchResponse{

		Data: employees,

		Page: req.Page,

		PageSize: req.PageSize,

		Total: total,

		TotalPages: totalPages,
	}, nil
}
