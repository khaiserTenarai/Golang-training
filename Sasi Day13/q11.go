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

type Employee struct {
	ID   int
	Name string
}

type Repository struct {
	queryTime time.Duration
}

func (r Repository) FindByName(ctx context.Context, name string) (Employee, error) {
	select {
	case <-time.After(r.queryTime):
		return Employee{ID: 1, Name: name}, nil
	case <-ctx.Done():
		return Employee{}, ctx.Err()
	}
}

type Service struct {
	repo Repository
}

func (s Service) Lookup(ctx context.Context, name string) (Employee, error) {
	if name == "" {
		return Employee{}, errors.New("name is required")
	}
	return s.repo.FindByName(ctx, name)
}

func handle(ctx context.Context, svc Service, name string) {
	emp, err := svc.Lookup(ctx, name)
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	fmt.Printf("found employee %d: %s\n", emp.ID, emp.Name)
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "Enter the employee name to search: ")
	timeoutMs, err1 := strconv.Atoi(ask(reader, "Enter the request timeout in milliseconds: "))
	queryMs, err2 := strconv.Atoi(ask(reader, "Enter how long the database query takes in milliseconds: "))
	if err1 != nil || err2 != nil {
		fmt.Println("both values must be numbers")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutMs)*time.Millisecond)
	defer cancel()

	svc := Service{repo: Repository{queryTime: time.Duration(queryMs) * time.Millisecond}}
	handle(ctx, svc, name)
}
