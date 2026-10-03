package main
import("q13-authentication-system/controller";"q13-authentication-system/database";"q13-authentication-system/repository";"q13-authentication-system/service";"q13-authentication-system/view")
func main(){db:=database.ConnectDB();defer db.Close();r:=repository.NewUserRepository(db);s:=service.NewAuthService(r);c:=controller.NewAuthController(s);view.NewAuthView(c).Start()}
