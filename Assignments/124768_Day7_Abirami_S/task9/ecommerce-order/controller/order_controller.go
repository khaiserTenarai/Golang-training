package controller
import("q09-ecommerce-order/model";"q09-ecommerce-order/service")
type OrderController interface{GetOrderDetails(int)([]model.OrderDetail,error)}
type OrderControllerImpl struct{service service.OrderService}
func NewOrderController(s service.OrderService)OrderController{return &OrderControllerImpl{s}}
func(c *OrderControllerImpl)GetOrderDetails(id int)([]model.OrderDetail,error){return c.service.GetOrderDetails(id)}
