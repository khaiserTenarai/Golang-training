package model

type Product struct {
	ID            int
	Name          string
	Description   string
	Price         float64
	Stock         int
	LowStockLimit int
}
