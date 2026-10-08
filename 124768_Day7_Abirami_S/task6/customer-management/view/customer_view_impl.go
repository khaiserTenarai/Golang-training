package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"customer-management/controller"
	"customer-management/model"
)

type CustomerViewImpl struct {
	controller controller.CustomerController
	reader     *bufio.Reader
}

func NewCustomerView(controller controller.CustomerController) CustomerView {
	return &CustomerViewImpl{
		controller: controller,
		reader:     bufio.NewReader(os.Stdin),
	}
}

func (v *CustomerViewImpl) Start() {
	for {
		fmt.Println("\n----- CUSTOMER MANAGEMENT -----")
		fmt.Println("1. Create Customer")
		fmt.Println("2. Get Customer")
		fmt.Println("3. Get All Customers")
		fmt.Println("4. Update Customer")
		fmt.Println("5. Delete Customer")
		fmt.Println("6. Search Customers")
		fmt.Println("7. Exit")

		fmt.Print("Enter choice: ")
		choice := v.readInt()

		switch choice {
		case 1:
			v.CreateCustomer()
		case 2:
			v.GetCustomer()
		case 3:
			v.GetAllCustomers()
		case 4:
			v.UpdateCustomer()
		case 5:
			v.DeleteCustomer()
		case 6:
			v.SearchCustomers()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (v *CustomerViewImpl) CreateCustomer() {
	var customer model.Customer

	fmt.Print("Enter Name: ")
	customer.Name = v.readString()

	fmt.Print("Enter Email: ")
	customer.Email = v.readString()

	fmt.Print("Enter Phone: ")
	customer.Phone = v.readString()

	fmt.Print("Enter City: ")
	customer.City = v.readString()

	err := v.controller.CreateCustomer(customer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Customer created successfully")
}

func (v *CustomerViewImpl) GetCustomer() {
	fmt.Print("Enter Customer ID: ")
	id := v.readInt()

	customer, err := v.controller.GetCustomer(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayCustomer(customer)
}

func (v *CustomerViewImpl) GetAllCustomers() {
	customers, err := v.controller.GetAllCustomers()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayCustomers(customers)
}

func (v *CustomerViewImpl) UpdateCustomer() {
	var customer model.Customer

	fmt.Print("Enter Customer ID: ")
	customer.ID = v.readInt()

	fmt.Print("Enter Name: ")
	customer.Name = v.readString()

	fmt.Print("Enter Email: ")
	customer.Email = v.readString()

	fmt.Print("Enter Phone: ")
	customer.Phone = v.readString()

	fmt.Print("Enter City: ")
	customer.City = v.readString()

	err := v.controller.UpdateCustomer(customer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Customer updated successfully")
}

func (v *CustomerViewImpl) DeleteCustomer() {
	fmt.Print("Enter Customer ID: ")
	id := v.readInt()

	err := v.controller.DeleteCustomer(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Customer deleted successfully")
}

func (v *CustomerViewImpl) SearchCustomers() {
	fmt.Println("\n----- SEARCH CUSTOMERS -----")

	fmt.Print("Enter Name (press Enter to skip): ")
	name := v.readString()

	fmt.Print("Enter Email (press Enter to skip): ")
	email := v.readString()

	fmt.Print("Enter City (press Enter to skip): ")
	city := v.readString()

	fmt.Print("Enter Page Number: ")
	page := v.readInt()

	fmt.Print("Enter Page Size: ")
	size := v.readInt()

	customers, err := v.controller.SearchCustomers(
		name,
		email,
		city,
		page,
		size,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayCustomers(customers)
}

func (v *CustomerViewImpl) DisplayCustomer(customer *model.Customer) {
	fmt.Println("\n----- CUSTOMER DETAILS -----")
	fmt.Println("ID    :", customer.ID)
	fmt.Println("Name  :", customer.Name)
	fmt.Println("Email :", customer.Email)
	fmt.Println("Phone :", customer.Phone)
	fmt.Println("City  :", customer.City)
}

func (v *CustomerViewImpl) DisplayCustomers(customers []model.Customer) {
	fmt.Println("\n----- CUSTOMER LIST -----")

	if len(customers) == 0 {
		fmt.Println("No customers found")
		return
	}

	for _, customer := range customers {
		fmt.Println("----------------------------")
		fmt.Println("ID    :", customer.ID)
		fmt.Println("Name  :", customer.Name)
		fmt.Println("Email :", customer.Email)
		fmt.Println("Phone :", customer.Phone)
		fmt.Println("City  :", customer.City)
	}
}

func (v *CustomerViewImpl) readString() string {
	value, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func (v *CustomerViewImpl) readInt() int {
	value := v.readString()
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return number
}
