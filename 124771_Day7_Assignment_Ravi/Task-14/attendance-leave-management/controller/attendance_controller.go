package controller

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "attendance-leave-management/model"
    "attendance-leave-management/service"
)

type AttendanceController struct {
    service service.AttendanceService
    reader  *bufio.Reader
}

func NewAttendanceController(
    service service.AttendanceService,
) *AttendanceController {
    return &AttendanceController{
        service: service,
        reader:  bufio.NewReader(os.Stdin),
    }
}

func (c *AttendanceController) Start() {
    for {
        fmt.Println()
        fmt.Println("========================================")
        fmt.Println("   ATTENDANCE & LEAVE MANAGEMENT")
        fmt.Println("========================================")
        fmt.Println("1. View Employees")
        fmt.Println("2. Employee Check-In")
        fmt.Println("3. Employee Check-Out")
        fmt.Println("4. Attendance Report")
        fmt.Println("5. Apply Leave")
        fmt.Println("6. View My Leaves")
        fmt.Println("7. View Pending Leaves")
        fmt.Println("8. Approve Leave")
        fmt.Println("9. Reject Leave")
        fmt.Println("10. Exit")
        fmt.Println("========================================")

        choice := c.readInt("Enter choice: ")

        switch choice {
        case 1:
            c.viewEmployees()
        case 2:
            c.checkIn()
        case 3:
            c.checkOut()
        case 4:
            c.attendanceReport()
        case 5:
            c.applyLeave()
        case 6:
            c.viewLeaves()
        case 7:
            c.viewPendingLeaves()
        case 8:
            c.approveLeave()
        case 9:
            c.rejectLeave()
        case 10:
            fmt.Println("Thank you for using the system.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func (c *AttendanceController) viewEmployees() {
    employees, err := c.service.GetEmployees()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("ID | Name | Email")
    fmt.Println("----------------------------------------")

    for _, employee := range employees {
        fmt.Printf(
            "%d | %s | %s
",
            employee.ID,
            employee.Name,
            employee.Email,
        )
    }
}

func (c *AttendanceController) checkIn() {
    employeeID := c.readInt("Enter employee ID: ")

    if err := c.service.CheckIn(employeeID); err != nil {
        fmt.Println("Check-in failed:", err)
        return
    }

    fmt.Println("Employee checked in successfully.")
}

func (c *AttendanceController) checkOut() {
    employeeID := c.readInt("Enter employee ID: ")

    if err := c.service.CheckOut(employeeID); err != nil {
        fmt.Println("Check-out failed:", err)
        return
    }

    fmt.Println("Employee checked out successfully.")
}

func (c *AttendanceController) attendanceReport() {
    employeeID := c.readInt("Enter employee ID: ")

    from := c.readDate("Enter from date (YYYY-MM-DD): ")
    to := c.readDate("Enter to date (YYYY-MM-DD): ")

    records, err := c.service.GetAttendanceReport(
        employeeID,
        from,
        to,
    )

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(records) == 0 {
        fmt.Println("No attendance records found.")
        return
    }

    fmt.Println()
    fmt.Println("========================================")
    fmt.Println("          ATTENDANCE REPORT")
    fmt.Println("========================================")

    for _, record := range records {
        checkIn := "-"
        checkOut := "-"

        if record.CheckIn != nil {
            checkIn = record.CheckIn.Format("15:04:05")
        }

        if record.CheckOut != nil {
            checkOut = record.CheckOut.Format("15:04:05")
        }

        fmt.Printf(
            "Date: %s | Check-In: %s | Check-Out: %s
",
            record.AttendanceDate.Format("2006-01-02"),
            checkIn,
            checkOut,
        )
    }
}

func (c *AttendanceController) applyLeave() {
    employeeID := c.readInt("Enter employee ID: ")
    leaveDate := c.readDate("Enter leave date (YYYY-MM-DD): ")
    reason := c.readString("Enter reason: ")

    leave := model.Leave{
        EmployeeID: employeeID,
        LeaveDate:  leaveDate,
        Reason:     reason,
    }

    if err := c.service.ApplyLeave(leave); err != nil {
        fmt.Println("Leave application failed:", err)
        return
    }

    fmt.Println("Leave applied successfully.")
    fmt.Println("Status: PENDING")
}

func (c *AttendanceController) viewLeaves() {
    employeeID := c.readInt("Enter employee ID: ")

    leaves, err := c.service.GetLeaves(employeeID)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    c.printLeaves(leaves)
}

func (c *AttendanceController) viewPendingLeaves() {
    leaves, err := c.service.GetAllPendingLeaves()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(leaves) == 0 {
        fmt.Println("No pending leaves.")
        return
    }

    c.printLeaves(leaves)
}

func (c *AttendanceController) approveLeave() {
    leaveID := c.readInt("Enter leave ID: ")

    if err := c.service.ApproveLeave(leaveID); err != nil {
        fmt.Println("Approval failed:", err)
        return
    }

    fmt.Println("Leave approved successfully.")
}

func (c *AttendanceController) rejectLeave() {
    leaveID := c.readInt("Enter leave ID: ")

    if err := c.service.RejectLeave(leaveID); err != nil {
        fmt.Println("Rejection failed:", err)
        return
    }

    fmt.Println("Leave rejected successfully.")
}

func (c *AttendanceController) printLeaves(leaves []model.Leave) {
    fmt.Println()
    fmt.Println("ID | Employee ID | Leave Date | Reason | Status")
    fmt.Println("---------------------------------------------------------------")

    for _, leave := range leaves {
        fmt.Printf(
            "%d | %d | %s | %s | %s
",
            leave.ID,
            leave.EmployeeID,
            leave.LeaveDate.Format("2006-01-02"),
            leave.Reason,
            leave.Status,
        )
    }
}

func (c *AttendanceController) readString(prompt string) string {
    fmt.Print(prompt)
    value, _ := c.reader.ReadString('
')
    return strings.TrimSpace(value)
}

func (c *AttendanceController) readInt(prompt string) int {
    for {
        value := c.readString(prompt)

        number, err := strconv.Atoi(value)
        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid number.")
    }
}

func (c *AttendanceController) readDate(prompt string) time.Time {
    for {
        value := c.readString(prompt)

        date, err := time.Parse("2006-01-02", value)
        if err == nil {
            return date
        }

        fmt.Println("Please use YYYY-MM-DD format.")
    }
}
