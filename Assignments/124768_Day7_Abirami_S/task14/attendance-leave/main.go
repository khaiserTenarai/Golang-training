package main
import("q14-attendance-leave/controller";"q14-attendance-leave/database";"q14-attendance-leave/repository";"q14-attendance-leave/service";"q14-attendance-leave/view")
func main(){db:=database.ConnectDB();defer db.Close();r:=repository.NewAttendanceRepository(db);s:=service.NewAttendanceService(r);c:=controller.NewAttendanceController(s);view.NewAttendanceView(c).Start()}
