package service
import("q14-attendance-leave/model";"q14-attendance-leave/repository")
type AttendanceService struct{repository repository.AttendanceRepository}
func NewAttendanceService(r repository.AttendanceRepository)*AttendanceService{return &AttendanceService{r}}
func(s *AttendanceService)CheckIn(id int)error{return s.repository.CheckIn(id)}
func(s *AttendanceService)CheckOut(id int)error{return s.repository.CheckOut(id)}
func(s *AttendanceService)GetAttendance(id int)([]model.Attendance,error){return s.repository.GetAttendance(id)}
func(s *AttendanceService)ApplyLeave(l model.Leave)error{return s.repository.ApplyLeave(l)}
func(s *AttendanceService)UpdateLeaveStatus(id int,x string)error{return s.repository.UpdateLeaveStatus(id,x)}
func(s *AttendanceService)GetLeaves(id int)([]model.Leave,error){return s.repository.GetLeaves(id)}
