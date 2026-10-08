package view
import("bufio";"fmt";"os";"q12-repository-service-architecture/controller";"q12-repository-service-architecture/model";"strconv";"strings")
type EmployeeView struct{c controller.EmployeeController;r *bufio.Reader}
func NewEmployeeView(c controller.EmployeeController)*EmployeeView{return &EmployeeView{c,bufio.NewReader(os.Stdin)}}
func(v *EmployeeView)Start(){for{fmt.Println("1.Create 2.Read 3.All 4.Update 5.Delete 6.Exit");fmt.Print("Choice: ");switch v.i(){case 1:v.create();case 2:v.read();case 3:v.all();case 4:v.update();case 5:fmt.Println(v.c.DeleteEmployee(v.i()));case 6:return}}}
func(v *EmployeeView)create(){var e model.Employee;fmt.Print("Name: ");e.Name=v.s();fmt.Print("Email: ");e.Email=v.s();fmt.Print("Salary: ");e.Salary=v.f();fmt.Println(v.c.CreateEmployee(e))}
func(v *EmployeeView)read(){e,x:=v.c.GetEmployee(v.i());if x!=nil{fmt.Println(x)}else{fmt.Println(e)}}
func(v *EmployeeView)all(){x,e:=v.c.GetAllEmployees();if e!=nil{fmt.Println(e);return};for _,a:=range x{fmt.Println(a)}}
func(v *EmployeeView)update(){var e model.Employee;fmt.Print("ID: ");e.ID=v.i();fmt.Print("Name: ");e.Name=v.s();fmt.Print("Email: ");e.Email=v.s();fmt.Print("Salary: ");e.Salary=v.f();fmt.Println(v.c.UpdateEmployee(e))}
func(v *EmployeeView)s()string{x,_:=v.r.ReadString('\n');return strings.TrimSpace(x)}
func(v *EmployeeView)i()int{x,_:=strconv.Atoi(v.s());return x}
func(v *EmployeeView)f()float64{x,_:=strconv.ParseFloat(v.s(),64);return x}
