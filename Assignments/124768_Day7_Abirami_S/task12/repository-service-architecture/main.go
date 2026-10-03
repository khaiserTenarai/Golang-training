package main
import("q12-repository-service-architecture/controller";"q12-repository-service-architecture/database";"q12-repository-service-architecture/repository";"q12-repository-service-architecture/service";"q12-repository-service-architecture/view")
func main(){db:=database.ConnectDB();defer db.Close();r:=repository.NewEmployeeRepository(db);s:=service.NewEmployeeService(r);c:=controller.NewEmployeeController(s);view.NewEmployeeView(c).Start()}
