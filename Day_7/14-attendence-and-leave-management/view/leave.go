package view
import (
	"fmt"

	"attendance_leave/model"
	"attendance_leave/utility"
)

type LeaveView interface {
	ShowMenu() int
	ReadLeave() model.Leave
	ReadID() int
	ReadEmployeeID() int
	DisplayLeaves(leaves []model.LeaveReport)
}

type LeaveViewImpl struct {
}

func NewLeaveView() LeaveView {
	return &LeaveViewImpl{}
}

func (v *LeaveViewImpl) ShowMenu() int {

	fmt.Println("\n========== Leave Management ==========")
	fmt.Println("1. Apply Leave")
	fmt.Println("2. Approve Leave")
	fmt.Println("3. Reject Leave")
	fmt.Println("4. Employee Leave Report")
	fmt.Println("5. All Leave Applications")
	fmt.Println("6. Back")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *LeaveViewImpl) ReadLeave() model.Leave {

	var leave model.Leave

	fmt.Println("\n---------- Apply Leave ----------")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&leave.EmployeeID)

	var fromDate string
	var toDate string

	fmt.Print("Enter From Date (YYYY-MM-DD): ")
	fmt.Scan(&fromDate)

	fmt.Print("Enter To Date (YYYY-MM-DD): ")
	fmt.Scan(&toDate)

	var err error

	leave.FromDate, err =
		utility.ParseDate(fromDate)

	if err != nil {
		fmt.Println("Invalid from date.")
		return leave
	}

	leave.ToDate, err =
		utility.ParseDate(toDate)

	if err != nil {
		fmt.Println("Invalid to date.")
		return leave
	}

	fmt.Print("Enter Reason: ")
	fmt.Scan(&leave.Reason)

	return leave
}

func (v *LeaveViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Leave ID: ")
	fmt.Scan(&id)

	return id
}

func (v *LeaveViewImpl) ReadEmployeeID() int {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	return id
}

func (v *LeaveViewImpl) DisplayLeaves(
	leaves []model.LeaveReport,
) {

	if len(leaves) == 0 {
		fmt.Println("No leave applications found.")
		return
	}

	fmt.Println("\n---------- Leave Applications ----------")

	for _, leave := range leaves {

		fmt.Println("Leave ID      :", leave.ID)
		fmt.Println("Employee ID   :", leave.EmployeeID)
		fmt.Println("Employee Name :", leave.EmployeeName)

		fmt.Println(
			"From Date     :",
			leave.FromDate.Format("2006-01-02"),
		)

		fmt.Println(
			"To Date       :",
			leave.ToDate.Format("2006-01-02"),
		)

		fmt.Println("Reason        :", leave.Reason)
		fmt.Println("Status        :", leave.Status)

		fmt.Println("----------------------------------------")
	}
}
