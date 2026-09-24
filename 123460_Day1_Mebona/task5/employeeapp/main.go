package main
import(
	"fmt"
	"employeeapp/employee"
)
func main(){
	details := employee.EmployeeName("John")
	fmt.Println(details)
}