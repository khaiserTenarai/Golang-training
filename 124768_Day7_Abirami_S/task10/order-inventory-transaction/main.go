package main
import("q10-order-inventory-transaction/controller";"q10-order-inventory-transaction/database";"q10-order-inventory-transaction/repository";"q10-order-inventory-transaction/service";"q10-order-inventory-transaction/view")
func main(){db:=database.ConnectDB();defer db.Close();r:=repository.NewOrderRepository(db);s:=service.NewOrderService(r);c:=controller.NewOrderController(s);view.NewOrderView(c).Start()}
