package main
import("q09-ecommerce-order/controller";"q09-ecommerce-order/database";"q09-ecommerce-order/repository";"q09-ecommerce-order/service";"q09-ecommerce-order/view")
func main(){db:=database.ConnectDB();defer db.Close();r:=repository.NewOrderRepository(db);s:=service.NewOrderService(r);c:=controller.NewOrderController(s);view.NewOrderView(c).Start()}
