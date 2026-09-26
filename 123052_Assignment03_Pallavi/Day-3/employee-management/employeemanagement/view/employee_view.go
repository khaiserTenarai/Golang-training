package view
import(
	"fmt"
 	"employee-management/employeemanagement/controller")
func ShowMessage() {
	message:=controller.GetMessage()
	fmt.Println(message)
	
}