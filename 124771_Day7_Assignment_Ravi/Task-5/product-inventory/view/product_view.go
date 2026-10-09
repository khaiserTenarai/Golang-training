package view

type CreateProductRequest struct {
	Name              string
	Description       string
	Price             float64
	StockQuantity     int
	LowStockThreshold int
}

type UpdateProductRequest struct {
	ID                int64
	Name              string
	Description       string
	Price             float64
	LowStockThreshold int
}

type StockRequest struct {
	ProductID int64
	Quantity  int
}
