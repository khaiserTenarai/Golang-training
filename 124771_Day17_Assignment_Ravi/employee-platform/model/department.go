package model

import "time"

type Department struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:120;not null;uniqueIndex" json:"name"`
	Description string     `gorm:"size:500" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	Employees   []Employee `json:"-"`
}
