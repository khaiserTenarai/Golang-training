package utility

import (
	"fmt"
	"strings"
)

var allowedSortColumns = map[string]bool{

	"id": true,

	"name": true,

	"department": true,

	"salary": true,
}

var allowedSortOrders = map[string]bool{

	"asc": true,

	"desc": true,
}

func ValidatePagination(
	page int,
	pageSize int,
) error {

	if page < 1 {

		return fmt.Errorf(
			"page must be greater than 0",
		)
	}

	if pageSize < 1 {

		return fmt.Errorf(
			"page size must be greater than 0",
		)
	}

	if pageSize > 100 {

		return fmt.Errorf(
			"page size cannot be greater than 100",
		)
	}

	return nil
}

func ValidateSort(
	sortBy string,
	sortOrder string,
) error {

	sortBy =
		strings.ToLower(sortBy)

	sortOrder =
		strings.ToLower(sortOrder)

	if !allowedSortColumns[sortBy] {

		return fmt.Errorf(
			"invalid sort field: %s",
			sortBy,
		)
	}

	if !allowedSortOrders[sortOrder] {

		return fmt.Errorf(
			"invalid sort order: %s",
			sortOrder,
		)
	}

	return nil
}

func ValidateSalary(
	salaryMin *float64,
	salaryMax *float64,
) error {

	if salaryMin != nil &&
		*salaryMin < 0 {

		return fmt.Errorf(
			"minimum salary cannot be negative",
		)
	}

	if salaryMax != nil &&
		*salaryMax < 0 {

		return fmt.Errorf(
			"maximum salary cannot be negative",
		)
	}

	if salaryMin != nil &&
		salaryMax != nil &&
		*salaryMin > *salaryMax {

		return fmt.Errorf(
			"minimum salary cannot be greater than maximum salary",
		)
	}

	return nil
}
