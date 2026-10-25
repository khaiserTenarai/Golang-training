package model

type Employee struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `gorm:"size:100;not null" json:"name"`
	Email      string  `gorm:"size:150;not null;uniqueIndex" json:"email"`
	Department string  `gorm:"size:100;not null" json:"department"`
	Salary     float64 `gorm:"not null" json:"salary"`
}
