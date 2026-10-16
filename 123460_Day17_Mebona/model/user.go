package model

// User is someone who can log in. Passwords are never stored - only a bcrypt hash.
type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"size:50;not null;uniqueIndex" json:"username"`
	PasswordHash string `gorm:"size:100;not null" json:"-"` // "-" = never sent in JSON
	Role         string `gorm:"size:20;not null" json:"role"`
}
