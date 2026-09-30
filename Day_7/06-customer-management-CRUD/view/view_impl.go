package view
import (
	"fmt"

	"cms/model"
)

type CustomerViewImpl struct {
}

func NewCustomerView() CustomerView {

	return &CustomerViewImpl{}
}

func (v *CustomerViewImpl) ShowMenu() int {

	fmt.Println("\n========== Customer Management System ==========")
	fmt.Println("1. Save Customer")
	fmt.Println("2. Find Customer")
	fmt.Println("3. Find All Customers")
	fmt.Println("4. Update Customer")
	fmt.Println("5. Delete Customer")
	fmt.Println("6. Search Customer")
	fmt.Println("7. Customer Pagination")
	fmt.Println("8. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *CustomerViewImpl) ReadCustomer() model.Customer {

	var customer model.Customer

	fmt.Println("\n---------- Enter Customer ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&customer.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&customer.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&customer.Email)

	fmt.Print("Enter Phone: ")
	fmt.Scan(&customer.Phone)

	return customer
}

func (v *CustomerViewImpl) ReadCustomerForUpdate() model.Customer {

	var customer model.Customer

	fmt.Println("\n---------- Update Customer ----------")

	fmt.Print("Enter Customer ID: ")
	fmt.Scan(&customer.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&customer.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&customer.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&customer.Email)

	fmt.Print("Enter Phone: ")
	fmt.Scan(&customer.Phone)

	return customer
}

func (v *CustomerViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Customer ID: ")
	fmt.Scan(&id)

	return id
}

func (v *CustomerViewImpl) ReadSearchKeyword() string {

	var keyword string

	fmt.Print("Enter search keyword: ")
	fmt.Scan(&keyword)

	return keyword
}

func (v *CustomerViewImpl) ReadPagination() (int, int) {

	var page int
	var limit int

	fmt.Print("Enter page number: ")
	fmt.Scan(&page)

	fmt.Print("Enter records per page: ")
	fmt.Scan(&limit)

	return page, limit
}

func (v *CustomerViewImpl) DisplayCustomer(
	customer model.Customer,
) {

	fmt.Println("\n---------- Customer ----------")

	fmt.Println("ID    :", customer.ID)
	fmt.Println("Name  :", customer.Name)
	fmt.Println("Age   :", customer.Age)
	fmt.Println("Email :", customer.Email)
	fmt.Println("Phone :", customer.Phone)
}

func (v *CustomerViewImpl) DisplayCustomers(
	customers []model.Customer,
) {

	if len(customers) == 0 {
		fmt.Println("No customers found.")
		return
	}

	fmt.Println("\n---------- Customers ----------")

	for _, customer := range customers {

		fmt.Println(
			customer.ID,
			customer.Name,
			customer.Age,
			customer.Email,
			customer.Phone,
		)
	}
}
