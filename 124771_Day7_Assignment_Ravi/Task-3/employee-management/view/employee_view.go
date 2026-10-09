package view

type EmployeeSearchRequest struct {
	Name       string
	Department string

	SalaryMin *float64
	SalaryMax *float64

	Page     int
	PageSize int

	SortBy    string
	SortOrder string
}
