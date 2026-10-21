package main

import "fmt"

// Task 3: Service and Repository separation using interfaces
// Using interfaces decouples the Business Logic (Service) from the Storage layer (Repository).
// This enables mock testing and seamless swapping of underlying database providers.

type DataRepository interface {
	FetchUser(id int) string
}

type MemoryRepository struct{}

func (m *MemoryRepository) FetchUser(id int) string {
	return fmt.Sprintf("User #%d from Memory Storage", id)
}

type UserService struct {
	Repo DataRepository
}

func (s *UserService) GetUserDetail(id int) string {
	return "Processed: " + s.Repo.FetchUser(id)
}

func main() {
	repo := &MemoryRepository{}
	svc := &UserService{Repo: repo}
	fmt.Println("Task 3 - Interface decoupling demo:", svc.GetUserDetail(42))
}
