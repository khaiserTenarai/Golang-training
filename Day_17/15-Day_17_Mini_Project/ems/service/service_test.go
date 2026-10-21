package service

import (
	"errors"
	"strings"
	"testing"

	"employee-management/model"
)

type fakeRepository struct {
	employees []model.Employee
}

func (f *fakeRepository) Create(e *model.Employee) error {
	e.ID = uint(len(f.employees) + 1)
	f.employees = append(f.employees, *e)
	return nil
}

func (f *fakeRepository) GetAll(
	filter model.EmployeeFilter,
	sort model.EmployeeSort,
	page, pageSize int,
) ([]model.Employee, int64, error) {
	var result []model.Employee

	for _, e := range f.employees {
		if filter.Name != "" &&
			!containsIgnoreCase(e.Name, filter.Name) {
			continue
		}
		if filter.Department != "" && e.Department != filter.Department {
			continue
		}
		result = append(result, e)
	}

	total := int64(len(result))
	start := (page - 1) * pageSize
	if start >= len(result) {
		return []model.Employee{}, total, nil
	}

	end := start + pageSize
	if end > len(result) {
		end = len(result)
	}
	return result[start:end], total, nil
}

func (f *fakeRepository) GetByID(id uint) (*model.Employee, error) {
	for i := range f.employees {
		if f.employees[i].ID == id {
			return &f.employees[i], nil
		}
	}
	return nil, errors.New("employee not found")
}

func (f *fakeRepository) Update(id uint, e *model.Employee) (*model.Employee, error) {
	for i := range f.employees {
		if f.employees[i].ID == id {
			e.ID = id
			f.employees[i] = *e
			return &f.employees[i], nil
		}
	}
	return nil, errors.New("employee not found")
}

func (f *fakeRepository) Delete(id uint) error {
	for i, e := range f.employees {
		if e.ID == id {
			f.employees = append(f.employees[:i], f.employees[i+1:]...)
			return nil
		}
	}
	return errors.New("employee not found")
}

// Small case-insensitive helper keeps this fake repository independent of SQL.
func containsIgnoreCase(value, sub string) bool {
	return len(sub) == 0 ||
		len(value) >= len(sub) &&
			strings.Contains(strings.ToLower(value), strings.ToLower(sub))
}

func TestCreateEmployeeValidatesInput(t *testing.T) {
	repo := &fakeRepository{}
	svc := NewEmployeeService(repo)

	err := svc.Create(&model.Employee{
		Name: "Priya", Email: "not-an-email",
		Department: "Finance", Salary: 50000,
	})
	if err == nil {
		t.Fatal("expected invalid email to be rejected")
	}
}

func TestFilteringAndPagination(t *testing.T) {
	repo := &fakeRepository{}
	svc := NewEmployeeService(repo)

	for _, e := range []model.Employee{
		{Name: "Amit", Email: "amit@example.com", Department: "IT", Salary: 50000},
		{Name: "Priya", Email: "priya@example.com", Department: "Finance", Salary: 65000},
		{Name: "Amit Kumar", Email: "amit.k@example.com", Department: "IT", Salary: 70000},
	} {
		copy := e
		if err := svc.Create(&copy); err != nil {
			t.Fatal(err)
		}
	}

	result, err := svc.GetAll(
		model.EmployeeFilter{Department: "IT"},
		model.EmployeeSort{Field: "id", Direction: "asc"},
		1, 1,
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.Total != 2 || len(result.Data) != 1 || result.TotalPages != 2 {
		t.Fatalf("unexpected pagination result: %+v", result)
	}
}

func TestRejectsInvalidPaginationAndSort(t *testing.T) {
	svc := NewEmployeeService(&fakeRepository{})

	_, err := svc.GetAll(model.EmployeeFilter{}, model.EmployeeSort{}, 0, 10)
	if err == nil {
		t.Fatal("expected invalid page to be rejected")
	}

	_, err = svc.GetAll(
		model.EmployeeFilter{},
		model.EmployeeSort{Field: "id; DROP TABLE employees", Direction: "asc"},
		1, 10,
	)
	if err == nil {
		t.Fatal("expected invalid sort field to be rejected")
	}
}
