package service

import (
	"ems/model"
	"errors"
	"testing"
)

// Mock repository satisfying the repository.EmployeeRepository interface
type MockEmployeeRepository struct {
	SaveFunc     func(employee model.Employee) error
	FindByIDFunc func(id int) (model.Employee, error)
	FindAllFunc  func() ([]model.Employee, error)
	UpdateFunc   func(employee model.Employee) error
	DeleteFunc   func(id int) error
}

func (m *MockEmployeeRepository) Save(e model.Employee) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(e)
	}
	return nil
}

func (m *MockEmployeeRepository) FindByID(id int) (model.Employee, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return model.Employee{}, nil
}

func (m *MockEmployeeRepository) FindAll() ([]model.Employee, error) {
	if m.FindAllFunc != nil {
		return m.FindAllFunc()
	}
	return nil, nil
}

func (m *MockEmployeeRepository) Update(e model.Employee) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(e)
	}
	return nil
}

func (m *MockEmployeeRepository) Delete(id int) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(id)
	}
	return nil
}

func TestEmployeeService_Save(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockEmployeeRepository{
			SaveFunc: func(e model.Employee) error {
				return nil
			},
		}

		service := NewEmployeeService(mockRepo)
		validEmp := model.Employee{Name: "Alice", Age: 25, Email: "alice@example.com"}

		err := service.Save(validEmp)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("Validation Failure - Invalid Age", func(t *testing.T) {
		mockRepo := &MockEmployeeRepository{}
		service := NewEmployeeService(mockRepo)
		invalidEmp := model.Employee{Name: "Alice", Age: 15, Email: "alice@example.com"}

		err := service.Save(invalidEmp)
		if err == nil {
			t.Error("expected validation error for underage employee, got nil")
		}
	})

	t.Run("Validation Failure - Invalid Email", func(t *testing.T) {
		mockRepo := &MockEmployeeRepository{}
		service := NewEmployeeService(mockRepo)
		invalidEmp := model.Employee{Name: "Alice", Age: 25, Email: "bademail"}

		err := service.Save(invalidEmp)
		if err == nil {
			t.Error("expected validation error for bad email, got nil")
		}
	})
}

func TestEmployeeService_FindByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockEmployeeRepository{
			FindByIDFunc: func(id int) (model.Employee, error) {
				return model.Employee{ID: id, Name: "Bob", Age: 30, Email: "bob@example.com"}, nil
			},
		}

		service := NewEmployeeService(mockRepo)
		emp, err := service.FindByID(1)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if emp.ID != 1 || emp.Name != "Bob" {
			t.Errorf("unexpected employee returned: %+v", emp)
		}
	})

	t.Run("Not Found", func(t *testing.T) {
		mockRepo := &MockEmployeeRepository{
			FindByIDFunc: func(id int) (model.Employee, error) {
				return model.Employee{}, errors.New("employee not found")
			},
		}

		service := NewEmployeeService(mockRepo)
		_, err := service.FindByID(99)

		if err == nil {
			t.Error("expected error for non-existent employee, got nil")
		}
	})
}
