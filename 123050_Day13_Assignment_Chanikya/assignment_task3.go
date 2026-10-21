package main

import "fmt"

// Task 3: Service and Repository Separation Using Interfaces
// Decouples domain logic from storage technology.
type DataRepository interface {
	FetchData() string
}

type MemoryRepository struct{}

func (r *MemoryRepository) FetchData() string {
	return "Data from Memory Slice"
}

type DataService struct {
	Repo DataRepository
}

func (s *DataService) ExecuteBusinessLogic() {
	data := s.Repo.FetchData()
	fmt.Printf("[Task 3] Service layer working via Interface: %s\n", data)
}

func main() {
	repo := &MemoryRepository{}
	svc := &DataService{Repo: repo}
	svc.ExecuteBusinessLogic()
}
