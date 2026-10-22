package test

import (
	"context"
	"errors"
	"testing"

	"employee-management/model"
	"employee-management/service"
)

type mockRepo struct {
	employees map[int64]model.Employee
	nextID    int64
}

func newMockRepo() *mockRepo {
	return &mockRepo{
		employees: make(map[int64]model.Employee),
		nextID:    1,
	}
}

func (m *mockRepo) Create(ctx context.Context, e *model.Employee) error {
	e.ID = m.nextID
	m.employees[e.ID] = *e
	m.nextID++
	return nil
}

func (m *mockRepo) GetAll(ctx context.Context, sortBy string) ([]model.Employee, error) {
	list := make([]model.Employee, 0, len(m.employees))
	for _, v := range m.employees {
		list = append(list, v)
	}
	return list, nil
}

func (m *mockRepo) GetByID(ctx context.Context, id int64) (*model.Employee, error) {
	e, ok := m.employees[id]
	if !ok {
		return nil, errors.New("employee not found")
	}
	return &e, nil
}

func (m *mockRepo) Update(ctx context.Context, id int64, e *model.Employee) (*model.Employee, error) {
	if _, ok := m.employees[id]; !ok {
		return nil, errors.New("employee not found")
	}
	e.ID = id
	m.employees[id] = *e
	return e, nil
}

func (m *mockRepo) Delete(ctx context.Context, id int64) error {
	if _, ok := m.employees[id]; !ok {
		return errors.New("employee not found")
	}
	delete(m.employees, id)
	return nil
}

func TestEmployeeService_Validation(t *testing.T) {
	repo := newMockRepo()
	svc := service.NewEmployeeService(repo)
	ctx := context.Background()

	t.Run("Create valid employee", func(t *testing.T) {
		emp := &model.Employee{
			Name:       "Jane Doe",
			Email:      "jane@example.com",
			Department: "Engineering",
			Salary:     90000,
		}
		err := svc.Create(ctx, emp)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if emp.ID != 1 {
			t.Errorf("expected ID 1, got %d", emp.ID)
		}
	})

	t.Run("Create employee invalid email", func(t *testing.T) {
		emp := &model.Employee{
			Name:       "Jane Doe",
			Email:      "invalid-email",
			Department: "Engineering",
			Salary:     90000,
		}
		err := svc.Create(ctx, emp)
		if err == nil {
			t.Fatal("expected email validation error, got nil")
		}
	})

	t.Run("Create employee negative salary", func(t *testing.T) {
		emp := &model.Employee{
			Name:       "Jane Doe",
			Email:      "jane@example.com",
			Department: "Engineering",
			Salary:     -100,
		}
		err := svc.Create(ctx, emp)
		if err == nil {
			t.Fatal("expected salary validation error, got nil")
		}
	})
}