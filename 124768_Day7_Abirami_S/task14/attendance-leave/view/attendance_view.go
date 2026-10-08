package view
import("bufio";"fmt";"os";"strconv";"strings";"time";"q14-attendance-leave/controller";"q14-attendance-leave/model")
type AttendanceView struct{c *controller.AttendanceController;r *bufio.Reader}
func NewAttendanceView(c *controller.AttendanceController)*AttendanceView{return &AttendanceView{c,bufio.NewReader(os.Stdin)}}
func(v *AttendanceView)Start(){for{fmt.Println("1.Check In 2.Check Out 3.Attendance 4.Apply Leave 5.Leave Status 6.Leaves 7.Exit");switch v.i(){case 1:fmt.Println(v.c.CheckIn(v.i()));case 2:fmt.Println(v.c.CheckOut(v.i()));case 3:x,e:=v.c.GetAttendance(v.i());if e!=nil{fmt.Println(e)}else{for _,a:=range x{fmt.Println(a)}};case 4:var l model.Leave;fmt.Print("Employee ID: ");l.EmployeeID=v.i();fmt.Print("Date YYYY-MM-DD: ");l.LeaveDate=v.date();fmt.Print("Reason: ");l.Reason=v.s();fmt.Println(v.c.ApplyLeave(l));case 5:fmt.Print("Leave ID: ");id:=v.i();fmt.Print("Status: ");fmt.Println(v.c.UpdateLeaveStatus(id,v.s()));case 6:x,e:=v.c.GetLeaves(v.i());if e!=nil{fmt.Println(e)}else{for _,l:=range x{fmt.Println(l)}};case 7:return}}}
func(v *AttendanceView)s()string{x,_:=v.r.ReadString('\n');return strings.TrimSpace(x)}
func(v *AttendanceView)i()int{x,_:=strconv.Atoi(v.s());return x}
func(v *AttendanceView)date()time.Time{x,_:=time.Parse("2006-01-02",v.s());return x}
