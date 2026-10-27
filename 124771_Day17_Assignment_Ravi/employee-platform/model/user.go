package model

import "time"

// Task 6: user records store a bcrypt hash, never the original password.
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:120;not null" json:"name"`
	Email        string    `gorm:"size:180;uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"size:20;not null;default:employee" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
