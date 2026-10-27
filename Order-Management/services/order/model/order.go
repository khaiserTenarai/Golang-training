package model

import "time"

type Order struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	CustomerID     uint      `json:"customerId" gorm:"not null;index:idx_orders_customer_created,priority:1"`
	ProductID      uint      `json:"productId" gorm:"not null;index"`
	Quantity       int       `json:"quantity" gorm:"not null;check:quantity > 0"`
	TotalAmount    float64   `json:"totalAmount" gorm:"not null"`
	Status         string    `json:"status" gorm:"not null;default:CREATED;index:idx_orders_status_created,priority:1"`
	StockReference string    `json:"-" gorm:"not null;default:'';size:80;index"`
	StockReleased  bool      `json:"stockReleased" gorm:"not null;default:false"`
	CreatedAt      time.Time `json:"createdAt" gorm:"index:idx_orders_customer_created,priority:2;index:idx_orders_status_created,priority:2"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (Order) TableName() string { return "orders" }
