package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"customer-management/model"
)

type CustomerView struct {
	reader *bufio.Reader
}

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewCustomerView() *CustomerView {

	return &CustomerView{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ==================================================
// MENU
// ==================================================

func (v *CustomerView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== CUSTOMER MANAGEMENT ==========")
	fmt.Println("1. Add Customer")
	fmt.Println("2. Find Customer By ID")
	fmt.Println("3. Find All Customers")
	fmt.Println("4. Update Customer")
	fmt.Println("5. Delete Customer")
	fmt.Println("6. Search Customers")
	fmt.Println("7. Customer Pagination")
	fmt.Println("8. Exit")
	fmt.Println("=========================================")
}

// ==================================================
// READ INTEGER
// ==================================================

func (v *CustomerView) ReadInt(
	message string,
) int {

	for {

		fmt.Print(message)

		input, err :=
			v.reader.ReadString('\n')

		if err != nil {
			fmt.Println(
				"Unable to read input.",
			)

			continue
		}

		input = strings.TrimSpace(input)

		value, err :=
			strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println(
			"Please enter a valid number.",
		)
	}
}

// ==================================================
// READ STRING
// ==================================================

func (v *CustomerView) ReadString(
	message string,
) string {

	fmt.Print(message)

	input, _ :=
		v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// ==================================================
// READ CUSTOMER
// ==================================================

func (v *CustomerView) ReadCustomer() model.Customer {

	return model.Customer{

		Name: v.ReadString(
			"Enter Customer Name: ",
		),

		Email: v.ReadString(
			"Enter Email: ",
		),

		Phone: v.ReadString(
			"Enter Phone: ",
		),

		City: v.ReadString(
			"Enter City: ",
		),
	}
}

// ==================================================
// READ CUSTOMER FOR UPDATE
// ==================================================

func (v *CustomerView) ReadCustomerForUpdate() model.Customer {

	return model.Customer{

		ID: v.ReadInt(
			"Enter Customer ID: ",
		),

		Name: v.ReadString(
			"Enter Customer Name: ",
		),

		Email: v.ReadString(
			"Enter Email: ",
		),

		Phone: v.ReadString(
			"Enter Phone: ",
		),

		City: v.ReadString(
			"Enter City: ",
		),
	}
}

// ==================================================
// SHOW ONE CUSTOMER
// ==================================================

func (v *CustomerView) ShowCustomer(
	customer model.Customer,
) {

	fmt.Println("--------------------------------")
	fmt.Println("ID    :", customer.ID)
	fmt.Println("Name  :", customer.Name)
	fmt.Println("Email :", customer.Email)
	fmt.Println("Phone :", customer.Phone)
	fmt.Println("City  :", customer.City)
	fmt.Println("--------------------------------")
}

// ==================================================
// SHOW CUSTOMERS
// ==================================================

func (v *CustomerView) ShowCustomers(
	customers []model.Customer,
) {

	if len(customers) == 0 {

		fmt.Println(
			"No customers found.",
		)

		return
	}

	for _, customer := range customers {

		v.ShowCustomer(customer)
	}
}

// ==================================================
// SHOW MESSAGE
// ==================================================

func (v *CustomerView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}
