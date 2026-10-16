package graphql

import (
	"employee-management/model"
	"employee-management/service"
	"gorm.io/gorm"
)

// These FAKE services keep data in memory, so the GraphQL tests need no database.

// ---------- fake employees ----------
type fakeEmployees struct {
	items      map[uint]*model.Employee
	nextID     uint
	lastFilter model.EmployeeFilter // remembers the last list request, so tests can check it
}

func newFakeEmployees() *fakeEmployees {
	f := &fakeEmployees{items: map[uint]*model.Employee{}, nextID: 2}
	f.items[1] = &model.Employee{
		ID: 1, Name: "Amit", Email: "amit@example.com", Salary: 75000, DepartmentID: 1,
		Department: &model.Department{ID: 1, Name: "IT"},
	}
	return f
}
func (f *fakeEmployees) Create(e *model.Employee) error {
	e.ID = f.nextID
	f.nextID++
	e.Department = &model.Department{ID: e.DepartmentID, Name: "IT"}
	cp := *e
	f.items[e.ID] = &cp
	return nil
}
func (f *fakeEmployees) List(flt model.EmployeeFilter) (*model.EmployeePage, error) {
	f.lastFilter = flt
	items := []model.Employee{}
	for _, e := range f.items {
		items = append(items, *e)
	}
	return &model.EmployeePage{Items: items, TotalCount: len(items), Page: flt.Page, PageSize: flt.PageSize, TotalPages: 1}, nil
}
func (f *fakeEmployees) ListByDepartment(id uint) ([]model.Employee, error) {
	items := []model.Employee{}
	for _, e := range f.items {
		if e.DepartmentID == id {
			items = append(items, *e)
		}
	}
	return items, nil
}
func (f *fakeEmployees) GetByID(id uint) (*model.Employee, error) {
	if e, ok := f.items[id]; ok {
		cp := *e
		return &cp, nil
	}
	return nil, gorm.ErrRecordNotFound
}
func (f *fakeEmployees) Update(id uint, e *model.Employee) (*model.Employee, error) {
	old, ok := f.items[id]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	old.Name, old.Email, old.Salary, old.DepartmentID = e.Name, e.Email, e.Salary, e.DepartmentID
	cp := *old
	return &cp, nil
}
func (f *fakeEmployees) Delete(id uint) error {
	if _, ok := f.items[id]; !ok {
		return gorm.ErrRecordNotFound
	}
	delete(f.items, id)
	return nil
}

// ---------- fake departments ----------
type fakeDepartments struct{}

func (fakeDepartments) Create(d *model.Department) error { d.ID = 2; return nil }
func (fakeDepartments) GetAll() ([]model.Department, error) {
	return []model.Department{{ID: 1, Name: "IT"}}, nil
}
func (fakeDepartments) GetByID(id uint) (*model.Department, error) {
	if id == 1 {
		return &model.Department{ID: 1, Name: "IT"}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

// ---------- fake login ----------
type fakeAuth struct{}

func (fakeAuth) Login(username, password string) (*service.LoginResult, error) {
	if username == "admin" && password == "admin123" {
		return &service.LoginResult{Token: "fake-token", Username: "admin", Role: "ADMIN"}, nil
	}
	return nil, service.ErrInvalidCredentials
}
func (fakeAuth) SeedDefaultUsers() error { return nil }
