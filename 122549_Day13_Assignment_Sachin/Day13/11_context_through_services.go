package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type contextKey string

const requestIDKey contextKey = "requestID"

type Employee struct {
	ID   int
	Name string
}

type Repository struct {
	queryTime time.Duration
}

func (r Repository) FindByID(ctx context.Context, id int) (Employee, error) {
	select {
	case <-time.After(r.queryTime): // the "query" finished
		return Employee{ID: id, Name: fmt.Sprintf("Employee %d", id)}, nil
	case <-ctx.Done(): // the caller gave up first
		return Employee{}, ctx.Err()
	}
}

//(middle layer)

type Service struct {
	repo Repository
}

func (s Service) GetEmployee(ctx context.Context, id int) (Employee, error) {
	reqID, _ := ctx.Value(requestIDKey).(string)
	fmt.Printf("[request %s] service: looking up employee %d\n", reqID, id)

	e, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return Employee{}, fmt.Errorf("get employee %d: %w", id, err)
	}
	return e, nil
}

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	id, err := strconv.Atoi(ask(reader, "Employee id to look up: "))
	if err != nil {
		fmt.Println("The id has to be a number.")
		return
	}
	queryFor, err := strconv.ParseFloat(ask(reader, "How many seconds does the database query take? "), 64)
	if err != nil || queryFor < 0 {
		fmt.Println("Please enter a number of seconds, 0 or more.")
		return
	}
	waitFor, err := strconv.ParseFloat(ask(reader, "How many seconds will the caller wait at most? "), 64)
	if err != nil || waitFor <= 0 {
		fmt.Println("Please enter a number of seconds above 0.")
		return
	}

	service := Service{repo: Repository{queryTime: time.Duration(queryFor * float64(time.Second))}}

	ctx := context.WithValue(context.Background(), requestIDKey, "req-001")
	ctx, cancel := context.WithTimeout(ctx, time.Duration(waitFor*float64(time.Second)))
	defer cancel()

	employee, err := service.GetEmployee(ctx, id)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("\nGave up waiting - the deadline passed before the query finished.")
		fmt.Println("Error:", err)
		return
	}
	if err != nil {
		fmt.Println("Something else went wrong:", err)
		return
	}
	fmt.Printf("\nFound: %+v\n", employee)
}
