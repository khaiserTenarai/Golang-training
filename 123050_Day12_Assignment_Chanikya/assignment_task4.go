package main

import "fmt"

// Task 4: Constructor-based Dependency Injection
// Constructors (e.g., NewService, NewRepository) guarantee that structs are fully 
// initialized with their required dependencies before execution, enforcing immutability and clarity.

type DBConn struct{ DSN string }

type AccountRepository struct {
	db *DBConn
}

func NewAccountRepository(db *DBConn) *AccountRepository {
	return &AccountRepository{db: db}
}

type AccountService struct {
	repo *AccountRepository
}

func NewAccountService(repo *AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func main() {
	db := &DBConn{DSN: "localhost:5432"}
	repo := NewAccountRepository(db)
	svc := NewAccountService(repo)
	fmt.Printf("Task 4 - Injected configuration successfully to target layer: %s\n", svc.repo.db.DSN)
}
