package service
import("q09-ecommerce-order/model";"q09-ecommerce-order/repository")
type OrderService interface{GetOrderDetails(int)([]model.OrderDetail,error)}
type OrderServiceImpl struct{repository repository.OrderRepository}
func NewOrderService(r repository.OrderRepository)OrderService{return &OrderServiceImpl{r}}
func(s *OrderServiceImpl)GetOrderDetails(id int)([]model.OrderDetail,error){return s.repository.GetOrderDetails(id)}
