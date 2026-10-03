package service

import (
	"testing"

	"ems/model"
)

type MockEmployeeRepository struct {
	SaveCalled bool
}

func (m *MockEmployeeRepository) Save(employee model.Employee) error {
	m.SaveCalled = true
	return nil
}

func (m *MockEmployeeRepository) FindByID(id int) (model.Employee, error) {
	return model.Employee{}, nil
}

func (m *MockEmployeeRepository) FindAll() ([]model.Employee, error) {
	return []model.Employee{}, nil
}

func (m *MockEmployeeRepository) Update(employee model.Employee) error {
	return nil
}

func (m *MockEmployeeRepository) Delete(id int) error {
	return nil
}

func TestSave_Success(t *testing.T) {

	mockRepo := &MockEmployeeRepository{}

	employeeService := NewEmployeeService(mockRepo)

	employee := model.Employee{
		Name:  "Ganesh",
		Age:   25,
		Email: "ganesh@gmail.com",
	}

	err := employeeService.Save(employee)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if !mockRepo.SaveCalled {
		t.Error("repository save not called")
	}
}
