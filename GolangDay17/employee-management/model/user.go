package model


type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"size:50;not null;uniqueIndex" json:"username"`
	PasswordHash string `gorm:"size:100;not null" json:"-"` 
	Role         string `gorm:"size:20;not null" json:"role"`
}
