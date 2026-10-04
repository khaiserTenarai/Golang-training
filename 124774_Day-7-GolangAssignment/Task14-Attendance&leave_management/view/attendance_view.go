package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"attendance-leave/model"
)

type AttendanceView struct {
	reader *bufio.Reader
}

func NewAttendanceView() *AttendanceView {

	return &AttendanceView{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (v *AttendanceView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== ATTENDANCE & LEAVE ==========")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Check In")
	fmt.Println("3. Check Out")
	fmt.Println("4. Attendance Report")
	fmt.Println("5. Apply Leave")
	fmt.Println("6. Approve Leave")
	fmt.Println("7. Reject Leave")
	fmt.Println("8. Leave Report")
	fmt.Println("9. Exit")
	fmt.Println("========================================")
}

func (v *AttendanceView) ReadString(
	message string,
) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (v *AttendanceView) ReadInt(
	message string,
) int {

	for {

		input := v.ReadString(message)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Enter a valid number.")
	}
}

func (v *AttendanceView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}

func (v *AttendanceView) ShowAttendance(
	list []model.Attendance,
) {

	if len(list) == 0 {
		fmt.Println("No attendance records.")
		return
	}

	for _, a := range list {

		fmt.Println("----------------------------")
		fmt.Println("Attendance ID:", a.ID)
		fmt.Println("Employee ID:", a.EmployeeID)
		fmt.Println("Check In:", a.CheckIn)
		fmt.Println("Check Out:", a.CheckOut)
	}
}

func (v *AttendanceView) ShowLeaves(
	list []model.Leave,
) {

	if len(list) == 0 {
		fmt.Println("No leave records.")
		return
	}

	for _, leave := range list {

		fmt.Println("----------------------------")
		fmt.Println("Leave ID:", leave.ID)
		fmt.Println("Employee ID:", leave.EmployeeID)
		fmt.Println("Reason:", leave.Reason)
		fmt.Println("Status:", leave.Status)
	}
}
