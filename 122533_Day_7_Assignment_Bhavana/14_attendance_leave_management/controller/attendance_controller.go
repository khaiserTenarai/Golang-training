package controller

import (
	"bufio"
	"context"
	"example.com/q14-attendance-leave/service"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AttendanceController struct {
	service service.AttendanceService
	reader  *bufio.Reader
}

func NewAttendanceController(s service.AttendanceService) *AttendanceController {
	return &AttendanceController{service: s, reader: bufio.NewReader(os.Stdin)}
}
func (c *AttendanceController) Start() {
	for {
		fmt.Println("\n1. Add Employee\n2. Check In\n3. Check Out\n4. Attendance Report\n5. Apply Leave\n6. Approve/Reject Leave\n7. Leave Report\n8. Exit")
		ch := c.i("Choice: ")
		ctx := context.Background()
		switch ch {
		case 1:
			c.done(c.service.AddEmployee(ctx, c.s("Employee name: ")))
		case 2:
			c.done(c.service.CheckIn(ctx, c.i64("Employee ID: ")))
		case 3:
			c.done(c.service.CheckOut(ctx, c.i64("Employee ID: ")))
		case 4:
			c.done(c.service.ShowAttendance(ctx))
		case 5:
			c.done(c.service.ApplyLeave(ctx, c.i64("Employee ID: "), c.s("Leave date (YYYY-MM-DD): "), c.s("Reason: ")))
		case 6:
			c.done(c.service.UpdateLeave(ctx, c.i64("Leave ID: "), c.s("Status (Approved/Rejected): ")))
		case 7:
			c.done(c.service.ShowLeaves(ctx))
		case 8:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (c *AttendanceController) done(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	} else {
		fmt.Println("Operation completed.")
	}
}
func (c *AttendanceController) s(p string) string {
	fmt.Print(p)
	v, _ := c.reader.ReadString('\n')
	return strings.TrimSpace(v)
}
func (c *AttendanceController) i(p string) int { v, _ := strconv.Atoi(c.s(p)); return v }
func (c *AttendanceController) i64(p string) int64 {
	v, _ := strconv.ParseInt(c.s(p), 10, 64)
	return v
}
