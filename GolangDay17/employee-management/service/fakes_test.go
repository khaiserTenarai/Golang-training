package service

import (
	"employee-management/model"
	"gorm.io/gorm"
)


type fakeEmployeeRepo struct {
	items      []model.Employee
	total      int64
	lastFilter model.EmployeeFilter
}

func (f *fakeEmployeeRepo) Create(e *model.Employee) error {
	e.ID = uint(len(f.items) + 1)
	f.items = append(f.items, *e)
	return nil
}
func (f *fakeEmployeeRepo) List(flt model.EmployeeFilter) ([]model.Employee, int64, error) {
	f.lastFilter = flt
	return f.items, f.total, nil
}
func (f *fakeEmployeeRepo) ListByDepartment(uint) ([]model.Employee, error) { return f.items, nil }
func (f *fakeEmployeeRepo) GetByID(id uint) (*model.Employee, error) {
	for i := range f.items {
		if f.items[i].ID == id {
			return &f.items[i], nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}
func (f *fakeEmployeeRepo) Update(uint, *model.Employee) (*model.Employee, error) {
	return nil, gorm.ErrRecordNotFound
}
func (f *fakeEmployeeRepo) Delete(uint) error { return gorm.ErrRecordNotFound }

type fakeDeptRepo struct{ exists map[uint]bool }

func (f *fakeDeptRepo) Create(*model.Department) error       { return nil }
func (f *fakeDeptRepo) GetAll() ([]model.Department, error) { return nil, nil }
func (f *fakeDeptRepo) GetByID(id uint) (*model.Department, error) {
	if f.exists[id] {
		return &model.Department{ID: id, Name: "IT"}, nil
	}
	return nil, gorm.ErrRecordNotFound
}
