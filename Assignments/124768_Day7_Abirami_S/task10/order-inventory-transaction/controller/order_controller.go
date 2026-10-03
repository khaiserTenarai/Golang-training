package controller
import("q10-order-inventory-transaction/model";"q10-order-inventory-transaction/service")
type OrderController interface{CreateOrder(model.Order)error}
type OrderControllerImpl struct{service service.OrderService}
func NewOrderController(s service.OrderService)OrderController{return &OrderControllerImpl{s}}
func(c *OrderControllerImpl)CreateOrder(o model.Order)error{return c.service.CreateOrder(o)}
