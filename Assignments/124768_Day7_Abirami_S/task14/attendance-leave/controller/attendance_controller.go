package controller
import("q14-attendance-leave/model";"q14-attendance-leave/service")
type AttendanceController struct{service *service.AttendanceService}
func NewAttendanceController(s *service.AttendanceService)*AttendanceController{return &AttendanceController{s}}
func(c *AttendanceController)CheckIn(id int)error{return c.service.CheckIn(id)}
func(c *AttendanceController)CheckOut(id int)error{return c.service.CheckOut(id)}
func(c *AttendanceController)GetAttendance(id int)([]model.Attendance,error){return c.service.GetAttendance(id)}
func(c *AttendanceController)ApplyLeave(l model.Leave)error{return c.service.ApplyLeave(l)}
func(c *AttendanceController)UpdateLeaveStatus(id int,x string)error{return c.service.UpdateLeaveStatus(id,x)}
func(c *AttendanceController)GetLeaves(id int)([]model.Leave,error){return c.service.GetLeaves(id)}
