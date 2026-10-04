package main

import (
	"context"
	"testing"
)

// MockRepository allows testing the Service layer without a real database
type MockRepository struct {
	users map[string]string
}

func (m *MockRepository) RegisterUser(username, hash string) error {
	m.users[username] = hash
	return nil
}

func (m *MockRepository) GetUserByUsername(username string) (*User, error) {
	hash, exists := m.users[username]
	if !exists {
		return nil, sql.ErrNoRows
	}
	return &User{ID: 1, Username: username, PasswordHash: hash}, nil
}

func (m *MockRepository) CreateEmployee(emp *Employee) error { return nil }
func (m *MockRepository) GetEmployees(limit, offset int) ([]Employee, error) { return nil, nil }
func (m *MockRepository) UpdateSalaryTx(ctx context.Context, empID int, newSalary float64) error { return nil }

func TestAuthService(t *testing.T) {
	mockRepo := &MockRepository{users: make(map[string]string)}
	svc := &AppService{repo: mockRepo}

	// Test Registration
	err := svc.Register("testuser", "password123")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	// Test Successful Authentication
	err = svc.Authenticate("testuser", "password123")
	if err != nil {
		t.Errorf("Authentication failed for valid credentials: %v", err)
	}

	// Test Failed Authentication
	err = svc.Authenticate("testuser", "wrongpassword")
	if err == nil {
		t.Errorf("Expected error for invalid password, got nil")
	}
}