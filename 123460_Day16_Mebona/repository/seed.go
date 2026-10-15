package repository

import (
	"time"

	"employee-management/model"

	"gorm.io/gorm"
)

// SeedSampleData creates 3 departments and, if the employees table is empty, 3 employees.
func SeedSampleData(db *gorm.DB) error {
	depts := []model.Department{
		{Name: "Engineering", Location: "Bengaluru"},
		{Name: "HR", Location: "Chennai"},
		{Name: "Finance", Location: "Mumbai"},
	}
	for i := range depts {
		if err := db.Where(model.Department{Name: depts[i].Name}).FirstOrCreate(&depts[i]).Error; err != nil {
			return err
		}
	}

	var count int64
	if err := db.Model(&model.Employee{}).Count(&count).Error; err != nil || count > 0 {
		return err
	}
	day := func(y int, m time.Month, d int) time.Time { return time.Date(y, m, d, 0, 0, 0, 0, time.UTC) }
	employees := []model.Employee{
		{Name: "Rajesh", Email: "rajesh@gmail.com", Salary: 60000, Active: true, DepartmentID: depts[0].ID, JoinedAt: day(2022, 1, 10)},
		{Name: "Amit", Email: "amit@gmail.com", Salary: 50000, Active: true, DepartmentID: depts[1].ID, JoinedAt: day(2023, 3, 15)},
		{Name: "Priya", Email: "priya@gmail.com", Salary: 75000, Active: true, DepartmentID: depts[0].ID, JoinedAt: day(2024, 6, 1)},
	}
	return db.Create(&employees).Error
}
