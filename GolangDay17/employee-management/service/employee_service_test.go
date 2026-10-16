package service

import (
	"employee-management/model"
	"testing"
)

func newTestService() (*fakeEmployeeRepo, EmployeeService) {
	repo := &fakeEmployeeRepo{}
	depts := &fakeDeptRepo{exists: map[uint]bool{1: true}}
	return repo, NewEmployeeService(repo, depts)
}

func TestCreateEmployee_Success(t *testing.T) {
	repo, svc := newTestService()
	e := model.Employee{Name: "Amit", Email: "amit@example.com", Salary: 50000, DepartmentID: 1}
	if err := svc.Create(&e); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if e.ID == 0 || len(repo.items) != 1 {
		t.Fatalf("employee was not saved")
	}
}

func TestCreateEmployee_Validation(t *testing.T) {
	_, svc := newTestService()
	tests := []struct {
		name string
		emp  model.Employee
	}{
		{"empty name", model.Employee{Name: "", Email: "a@b.com", Salary: 1000, DepartmentID: 1}},
		{"name too short", model.Employee{Name: "A", Email: "a@b.com", Salary: 1000, DepartmentID: 1}},
		{"bad email", model.Employee{Name: "Amit", Email: "not-an-email", Salary: 1000, DepartmentID: 1}},
		{"negative salary", model.Employee{Name: "Amit", Email: "a@b.com", Salary: -5, DepartmentID: 1}},
		{"salary too large", model.Employee{Name: "Amit", Email: "a@b.com", Salary: 99999999, DepartmentID: 1}},
		{"missing department", model.Employee{Name: "Amit", Email: "a@b.com", Salary: 1000}},
		{"unknown department", model.Employee{Name: "Amit", Email: "a@b.com", Salary: 1000, DepartmentID: 99}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := tc.emp
			err := svc.Create(&e)
			if !IsValidation(err) {
				t.Fatalf("expected a validation error, got %v", err)
			}
		})
	}
}

func TestList_Defaults(t *testing.T) {
	repo, svc := newTestService()
	page, err := svc.List(model.EmployeeFilter{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if page.Page != 1 || page.PageSize != 10 {
		t.Fatalf("expected page 1 / size 10, got %d / %d", page.Page, page.PageSize)
	}
	if repo.lastFilter.SortBy != "id" || repo.lastFilter.SortOrder != "asc" {
		t.Fatalf("expected default sort id asc, got %s %s", repo.lastFilter.SortBy, repo.lastFilter.SortOrder)
	}
}

func TestList_PaginationMath(t *testing.T) {
	repo, svc := newTestService()
	repo.total = 25 // 25 employees, 10 per page -> 3 pages

	first, _ := svc.List(model.EmployeeFilter{Page: 1, PageSize: 10})
	if first.TotalPages != 3 || !first.HasNext || first.TotalCount != 25 {
		t.Fatalf("page 1 wrong: %+v", first)
	}
	last, _ := svc.List(model.EmployeeFilter{Page: 3, PageSize: 10})
	if last.HasNext {
		t.Fatalf("page 3 should be the last page")
	}
}

func TestList_RejectsBadInput(t *testing.T) {
	_, svc := newTestService()
	lo, hi := 5000.0, 1000.0
	tests := []struct {
		name string
		f    model.EmployeeFilter
	}{
		{"negative page", model.EmployeeFilter{Page: -1}},
		{"page size too big", model.EmployeeFilter{PageSize: 500}},
		{"unknown sort field", model.EmployeeFilter{SortBy: "password"}},
		{"unknown sort order", model.EmployeeFilter{SortOrder: "sideways"}},
		{"min greater than max", model.EmployeeFilter{MinSalary: &lo, MaxSalary: &hi}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := svc.List(tc.f); !IsValidation(err) {
				t.Fatalf("expected a validation error, got %v", err)
			}
		})
	}
}
