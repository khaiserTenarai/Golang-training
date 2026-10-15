package graphql

import (
	"sort"
	"strings"
	"sync"
	"time"

	"employee-management/model"
	"employee-management/repository"
)

// In-memory repositories so the whole GraphQL stack can be tested without PostgreSQL.

type fakeDeptRepo struct{ items []*model.Department }

func newFakeDeptRepo() *fakeDeptRepo {
	return &fakeDeptRepo{items: []*model.Department{
		{ID: 1, Name: "Engineering", Location: "Bengaluru"},
		{ID: 2, Name: "HR", Location: "Chennai"},
		{ID: 3, Name: "Finance", Location: "Mumbai"},
	}}
}

func (f *fakeDeptRepo) List() ([]*model.Department, error) { return f.items, nil }

func (f *fakeDeptRepo) GetByID(id uint) (*model.Department, error) {
	for _, d := range f.items {
		if d.ID == id {
			cp := *d
			return &cp, nil
		}
	}
	return nil, repository.ErrNotFound
}

type fakeEmpRepo struct {
	mu    sync.Mutex
	next  uint
	data  map[uint]model.Employee
	depts *fakeDeptRepo
}

func newFakeEmpRepo(depts *fakeDeptRepo) *fakeEmpRepo {
	f := &fakeEmpRepo{next: 1, data: map[uint]model.Employee{}, depts: depts}
	day := func(y int, m time.Month, d int) time.Time { return time.Date(y, m, d, 0, 0, 0, 0, time.UTC) }
	_ = f.Create(&model.Employee{Name: "Rajesh", Email: "rajesh@gmail.com", Salary: 60000, Active: true, DepartmentID: 1, JoinedAt: day(2022, 1, 10)})
	_ = f.Create(&model.Employee{Name: "Amit", Email: "amit@gmail.com", Salary: 50000, Active: true, DepartmentID: 2, JoinedAt: day(2023, 3, 15)})
	_ = f.Create(&model.Employee{Name: "Priya", Email: "priya@gmail.com", Salary: 75000, Active: true, DepartmentID: 1, JoinedAt: day(2024, 6, 1)})
	return f
}

// withDept mimics GORM's Preload("Department").
func (f *fakeEmpRepo) withDept(e model.Employee) *model.Employee {
	if d, err := f.depts.GetByID(e.DepartmentID); err == nil {
		e.Department = *d
	}
	return &e
}

func (f *fakeEmpRepo) Create(e *model.Employee) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	e.ID = f.next
	f.next++
	f.data[e.ID] = *e
	return nil
}

func (f *fakeEmpRepo) GetByID(id uint) (*model.Employee, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	e, ok := f.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return f.withDept(e), nil
}

func (f *fakeEmpRepo) GetByEmail(email string) (*model.Employee, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, e := range f.data {
		if strings.EqualFold(e.Email, email) {
			return f.withDept(e), nil
		}
	}
	return nil, repository.ErrNotFound
}

func (f *fakeEmpRepo) List(flt model.EmployeeFilter) ([]*model.Employee, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	search := strings.ToLower(flt.Search)
	var out []*model.Employee
	for _, e := range f.data {
		if flt.DepartmentID != nil && e.DepartmentID != *flt.DepartmentID {
			continue
		}
		if flt.JoinedAfter != nil && !e.JoinedAt.After(*flt.JoinedAfter) {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(e.Name), search) &&
			!strings.Contains(strings.ToLower(e.Email), search) {
			continue
		}
		out = append(out, f.withDept(e))
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	if flt.Offset >= len(out) {
		return []*model.Employee{}, nil
	}
	out = out[flt.Offset:]
	if flt.Limit > 0 && flt.Limit < len(out) {
		out = out[:flt.Limit]
	}
	return out, nil
}

func (f *fakeEmpRepo) Update(e *model.Employee) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.data[e.ID]; !ok {
		return repository.ErrNotFound
	}
	f.data[e.ID] = *e
	return nil
}

func (f *fakeEmpRepo) Delete(id uint) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.data[id]; !ok {
		return repository.ErrNotFound
	}
	delete(f.data, id)
	return nil
}
