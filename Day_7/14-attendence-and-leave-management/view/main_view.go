package view
import "fmt"

type MainView interface {
	ShowMenu() int
}

type MainViewImpl struct {
}

func NewMainView() MainView {
	return &MainViewImpl{}
}

func (v *MainViewImpl) ShowMenu() int {

	fmt.Println("\n==========================================")
	fmt.Println("     ATTENDANCE & LEAVE MANAGEMENT")
	fmt.Println("==========================================")
	fmt.Println("1. Employee Management")
	fmt.Println("2. Attendance Management")
	fmt.Println("3. Leave Management")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}
