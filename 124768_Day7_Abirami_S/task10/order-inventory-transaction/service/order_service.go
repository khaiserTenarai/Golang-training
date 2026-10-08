package service
import("q10-order-inventory-transaction/model";"q10-order-inventory-transaction/repository")
type OrderService interface{CreateOrder(model.Order)error}
type OrderServiceImpl struct{repository repository.OrderRepository}
func NewOrderService(r repository.OrderRepository)OrderService{return &OrderServiceImpl{r}}
func(s *OrderServiceImpl)CreateOrder(o model.Order)error{return s.repository.CreateOrder(o)}
