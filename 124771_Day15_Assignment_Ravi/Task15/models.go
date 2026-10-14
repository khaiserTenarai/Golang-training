package main

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	OwnerID   string `json:"owner_id"`
}

var (
	users      = make(map[string]User)
	tasks      = make(map[int]Task)
	nextTaskID = 1
)
