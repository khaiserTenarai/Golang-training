package repository

type OrderRepository interface {
	CreateOrder(
		customerName string,
		productID int,
		quantity int,
	) error
}
