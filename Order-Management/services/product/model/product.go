package model

import "time"

// Product represents a product stored in PostgreSQL.
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null;size:150"`
	Description string    `json:"description" gorm:"type:text"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null;default:0"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName specifies the PostgreSQL table name.
func (Product) TableName() string {
	return "products"
}
