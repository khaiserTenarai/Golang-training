package database

import (
	"fmt"

	"day17-employee-platform/config"
	"day17-employee-platform/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Department{},
		&model.Employee{},
	)
}

func Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.Department{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	departments := []model.Department{
		{Name: "IT"},
		{Name: "HR"},
		{Name: "Finance"},
	}

	if err := db.Create(&departments).Error; err != nil {
		return err
	}

	employees := []model.Employee{
		{Name: "Alice", Email: "alice@example.com", Salary: 50000, DepartmentID: departments[0].ID},
		{Name: "Bob", Email: "bob@example.com", Salary: 45000, DepartmentID: departments[1].ID},
		{Name: "Charlie", Email: "charlie@example.com", Salary: 55000, DepartmentID: departments[2].ID},
	}

	return db.Create(&employees).Error
}
