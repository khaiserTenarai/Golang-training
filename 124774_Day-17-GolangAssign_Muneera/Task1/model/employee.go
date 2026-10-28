package model

type Employee struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name" gorm:"not null"`
	Email      string  `json:"email" gorm:"unique;not null"`
	Department string  `json:"department" gorm:"not null"`
	Salary     float64 `json:"salary" gorm:"not null"`
}
