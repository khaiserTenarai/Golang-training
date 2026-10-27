package model

import "time"

// StockOperation is an idempotency ledger for stock reservations/releases.
type StockOperation struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	OrderReference string    `json:"orderReference" gorm:"not null;uniqueIndex:idx_stock_order_action,priority:1;size:80"`
	Action         string    `json:"action" gorm:"not null;uniqueIndex:idx_stock_order_action,priority:2;size:16"`
	ProductID      uint      `json:"productId" gorm:"not null;index"`
	Quantity       int       `json:"quantity" gorm:"not null"`
	UnitPrice      float64   `json:"unitPrice" gorm:"not null;default:0"`
	CreatedAt      time.Time `json:"createdAt"`
}

func (StockOperation) TableName() string { return "product_stock_operations" }
