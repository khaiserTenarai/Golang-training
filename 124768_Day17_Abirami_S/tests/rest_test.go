package tests

import (
	"testing"

	"day17-employee-platform/config"
	"day17-employee-platform/database"
	"day17-employee-platform/model"
	"day17-employee-platform/repository"
	"day17-employee-platform/service"
)

func TestEmployeeService(t *testing.T) {
	cfg := config.Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "Root",
		DBName:     "day17_employee_db",
	}

	db, err := database.Connect(cfg)
	if err != nil {
		t.Skip("PostgreSQL is not available")
	}

	if err := database.Migrate(db); err != nil {
		t.Skip("PostgreSQL database is not available")
	}

	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)

	_, _, _, err = svc.GetAll(
		model.EmployeeFilter{},
		model.SortOption{Field: "name", Order: "asc"},
		1,
		5,
	)
	if err != nil {
		t.Fatal(err)
	}
}
