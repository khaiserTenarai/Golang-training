package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"attendance_leave_management/controller"
	"attendance_leave_management/model"
	"github.com/jackc/pgx/v5"
)

func Start(conn *pgx.Conn) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("===== ATTENDANCE & LEAVE MANAGEMENT =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Display Employees")
		fmt.Println("3. Employee Check-In")
		fmt.Println("4. Employee Check-Out")
		fmt.Println("5. Attendance Report")
		fmt.Println("6. Apply Leave")
		fmt.Println("7. Display Leave Applications")
		fmt.Println("8. Approve Leave")
		fmt.Println("9. Reject Leave")
		fmt.Println("10. Exit")

		choice := readInt(reader, "Enter choice: ")

		switch choice {
		case 1:
			addEmployee(conn, reader)
		case 2:
			displayEmployees(conn)
		case 3:
			checkIn(conn, reader)
		case 4:
			checkOut(conn, reader)
		case 5:
			attendanceReport(conn)
		case 6:
			applyLeave(conn, reader)
		case 7:
			displayLeaves(conn)
		case 8:
			approveLeave(conn, reader)
		case 9:
			rejectLeave(conn, reader)
		case 10:
			fmt.Println("Program ended.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addEmployee(conn *pgx.Conn, r *bufio.Reader) {
	e := model.Employee{
		Name:       readString(r, "Enter name: "),
		Email:      readString(r, "Enter email: "),
		Department: readString(r, "Enter department: "),
	}
	if err := controller.AddEmployee(conn, e); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Employee added successfully.")
}

func displayEmployees(conn *pgx.Conn) {
	employees, err := controller.GetEmployees(conn)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== EMPLOYEES =====")
	fmt.Println("ID | Name | Email | Department")
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, e := range employees {
		fmt.Printf("%d | %s | %s | %s\n", e.ID, e.Name, e.Email, e.Department)
	}
}

func checkIn(conn *pgx.Conn, r *bufio.Reader) {
	id := readInt(r, "Enter employee ID: ")
	if err := controller.CheckIn(conn, id); err != nil {
		fmt.Println("Check-in failed:", err)
		return
	}
	fmt.Println("Employee checked in successfully.")
}

func checkOut(conn *pgx.Conn, r *bufio.Reader) {
	id := readInt(r, "Enter employee ID: ")
	if err := controller.CheckOut(conn, id); err != nil {
		fmt.Println("Check-out failed:", err)
		return
	}
	fmt.Println("Employee checked out successfully.")
}

func attendanceReport(conn *pgx.Conn) {
	report, err := controller.GetAttendanceReport(conn)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== ATTENDANCE REPORT =====")
	fmt.Println("ID | Employee | Date | Check-In | Check-Out")
	if len(report) == 0 {
		fmt.Println("No attendance records found.")
		return
	}

	for _, a := range report {
		in, out := "-", "-"
		if a.CheckIn != nil {
			in = a.CheckIn.Format("2006-01-02 15:04:05")
		}
		if a.CheckOut != nil {
			out = a.CheckOut.Format("2006-01-02 15:04:05")
		}

		fmt.Printf("%d | %s | %s | %s | %s\n",
			a.ID, a.EmployeeName, a.AttendanceDate.Format("2006-01-02"), in, out)
	}
}

func applyLeave(conn *pgx.Conn, r *bufio.Reader) {
	id := readInt(r, "Enter employee ID: ")
	typ := readString(r, "Enter leave type: ")
	startText := readString(r, "Enter start date (YYYY-MM-DD): ")
	endText := readString(r, "Enter end date (YYYY-MM-DD): ")
	reason := readString(r, "Enter reason: ")

	start, err := controller.ParseDate(startText)
	if err != nil {
		fmt.Println("Invalid start date. Use YYYY-MM-DD.")
		return
	}

	end, err := controller.ParseDate(endText)
	if err != nil {
		fmt.Println("Invalid end date. Use YYYY-MM-DD.")
		return
	}

	err = controller.ApplyLeave(conn, model.Leave{
		EmployeeID: id, LeaveType: typ, StartDate: start, EndDate: end, Reason: reason,
	})
	if err != nil {
		fmt.Println("Leave application failed:", err)
		return
	}

	fmt.Println("Leave application submitted successfully.")
	fmt.Println("Status: Pending")
}

func displayLeaves(conn *pgx.Conn) {
	leaves, err := controller.GetLeaves(conn)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== LEAVE APPLICATIONS =====")
	fmt.Println("ID | Employee | Type | Start | End | Reason | Status")
	if len(leaves) == 0 {
		fmt.Println("No leave applications found.")
		return
	}

	for _, l := range leaves {
		fmt.Printf("%d | %s | %s | %s | %s | %s | %s\n",
			l.ID, l.EmployeeName, l.LeaveType,
			l.StartDate.Format("2006-01-02"), l.EndDate.Format("2006-01-02"),
			l.Reason, l.Status)
	}
}

func approveLeave(conn *pgx.Conn, r *bufio.Reader) {
	id := readInt(r, "Enter leave ID to approve: ")
	if err := controller.ApproveLeave(conn, id); err != nil {
		fmt.Println("Approval failed:", err)
		return
	}
	fmt.Println("Leave approved successfully.")
}

func rejectLeave(conn *pgx.Conn, r *bufio.Reader) {
	id := readInt(r, "Enter leave ID to reject: ")
	if err := controller.RejectLeave(conn, id); err != nil {
		fmt.Println("Rejection failed:", err)
		return
	}
	fmt.Println("Leave rejected successfully.")
}

func readString(r *bufio.Reader, message string) string {
	for {
		fmt.Print(message)
		value, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}
		value = strings.TrimSpace(value)
		if value == "" {
			fmt.Println("Input cannot be empty.")
			continue
		}
		return value
	}
}

func readInt(r *bufio.Reader, message string) int {
	for {
		value := readString(r, message)
		n, err := strconv.Atoi(value)
		if err == nil {
			return n
		}
		fmt.Println("Please enter a valid number.")
	}
}
