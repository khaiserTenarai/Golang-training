package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management/controller"
)

type SalaryViewImpl struct {
	controller controller.SalaryController
	reader     *bufio.Reader
}

func NewSalaryView(controller controller.SalaryController) SalaryView {
	return &SalaryViewImpl{
		controller: controller,
		reader:     bufio.NewReader(os.Stdin),
	}
}

func (v *SalaryViewImpl) UpdateSalary() {
	fmt.Println("\n----- UPDATE EMPLOYEE SALARY -----")

	fmt.Print("Enter Employee ID: ")
	employeeID := v.readInt()

	fmt.Print("Enter New Salary: ")
	newSalary := v.readFloat()

	err := v.controller.UpdateSalary(employeeID, newSalary)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Salary updated successfully")
}

func (v *SalaryViewImpl) readString() string {
	value, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func (v *SalaryViewImpl) readInt() int {
	value := v.readString()
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return number
}

func (v *SalaryViewImpl) readFloat() float64 {
	value := v.readString()
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return number
}
