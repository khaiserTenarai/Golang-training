package service

type OrderService interface {
	CreateOrder(
		customerName string,
		productID int,
		quantity int,
	) error
}
