package service_test

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"employee-api/model"
	"employee-api/repository"
	"employee-api/service"
)

func newService(searchEnabled bool) *service.EmployeeService {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	opts := service.Options{RequestTimeout: time.Second, DefaultPageSize: 10, MaxPageSize: 100, SearchEnabled: searchEnabled}
	return service.NewEmployeeService(repository.NewMemoryEmployeeRepository(), opts, logger)
}

func input(email string) model.EmployeeInput {
	return model.EmployeeInput{Name: "Riya", Email: email, Department: "IT", Salary: 50000}
}

func TestCreateAndGet(t *testing.T) {
	svc := newService(true)
	created, err := svc.Create(context.Background(), input("Riya@Example.com "))
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != 1 || created.Email != "riya@example.com" {
		t.Fatalf("unexpected employee: %+v", created)
	}
	got, err := svc.Get(context.Background(), created.ID)
	if err != nil || got.Name != "Riya" {
		t.Fatalf("get failed: %v %+v", err, got)
	}
	if _, err := svc.Get(context.Background(), 99); !errors.Is(err, model.ErrNotFound) {
		t.Fatalf("want ErrNotFound, got %v", err)
	}
}

func TestValidationAndConflict(t *testing.T) {
	svc := newService(true)
	var ve *model.ValidationError
	if _, err := svc.Create(context.Background(), model.EmployeeInput{Name: "R", Email: "nope", Salary: -1}); !errors.As(err, &ve) {
		t.Fatalf("want ValidationError, got %v", err)
	}
	if len(ve.Fields) != 4 {
		t.Fatalf("want 4 invalid fields, got %v", ve.Fields)
	}

	if _, err := svc.Create(context.Background(), input("a@example.com")); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.Create(context.Background(), input("a@example.com")); !errors.Is(err, model.ErrConflict) {
		t.Fatalf("want ErrConflict, got %v", err)
	}
}

func TestSearchFeatureFlag(t *testing.T) {
	enabled, disabled := newService(true), newService(false)
	_, _ = enabled.Create(context.Background(), input("a@example.com"))

	res, err := enabled.List(context.Background(), service.ListQuery{Search: "riya"})
	if err != nil || res.Total != 1 {
		t.Fatalf("search should work: %v %+v", err, res)
	}
	if _, err := disabled.List(context.Background(), service.ListQuery{Search: "riya"}); !errors.Is(err, model.ErrFeatureDisabled) {
		t.Fatalf("want ErrFeatureDisabled, got %v", err)
	}
	if _, err := disabled.List(context.Background(), service.ListQuery{}); err != nil {
		t.Fatalf("listing without search must still work: %v", err)
	}
}

func TestPagination(t *testing.T) {
	svc := newService(true)
	for i := 0; i < 25; i++ {
		_, _ = svc.Create(context.Background(), input(string(rune('a'+i))+"@example.com"))
	}
	res, err := svc.List(context.Background(), service.ListQuery{Page: 3, PageSize: 10})
	if err != nil || len(res.Items) != 5 || res.Total != 25 || res.TotalPages != 3 {
		t.Fatalf("unexpected page: %v %+v", err, res)
	}
	var ve *model.ValidationError
	if _, err := svc.List(context.Background(), service.ListQuery{PageSize: 1000}); !errors.As(err, &ve) {
		t.Fatalf("pageSize above max must be rejected, got %v", err)
	}
}

func TestContextCancellationReachesTheRepository(t *testing.T) {
	svc := newService(true)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := svc.Create(ctx, input("a@example.com")); !errors.Is(err, context.Canceled) {
		t.Fatalf("want context.Canceled, got %v", err)
	}
}
