package view

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"task06_customer_management/models"
)

type CustomerView struct {
	scanner *bufio.Scanner
}

func NewCustomerView() *CustomerView {
	return &CustomerView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *CustomerView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *CustomerView) ShowMenu() int {
	fmt.Println("\n===== CUSTOMER MANAGEMENT =====")
	fmt.Println("1. Add Customer")
	fmt.Println("2. List Customers (Paginated)")
	fmt.Println("3. Get Customer by ID")
	fmt.Println("4. Update Customer")
	fmt.Println("5. Delete Customer")
	fmt.Println("6. Search Customers")
	fmt.Println("7. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *CustomerView) GetCustomerInput() models.Customer {
	var c models.Customer
	fmt.Print("Enter Name: ")
	c.Name = v.readLine()
	fmt.Print("Enter Email: ")
	c.Email = v.readLine()
	fmt.Print("Enter Phone: ")
	c.Phone = v.readLine()
	fmt.Print("Enter City: ")
	c.City = v.readLine()
	return c
}

func (v *CustomerView) GetID() int {
	fmt.Print("Enter Customer ID: ")
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *CustomerView) GetPaginationInput() (int, int) {
	fmt.Print("Page number: ")
	page, _ := strconv.Atoi(v.readLine())
	fmt.Print("Page size: ")
	pageSize, _ := strconv.Atoi(v.readLine())
	return page, pageSize
}

func (v *CustomerView) GetSearchKeyword() string {
	fmt.Print("Enter search keyword: ")
	return v.readLine()
}

func (v *CustomerView) GetUpdateInput(existing models.Customer) models.Customer {
	fmt.Printf("Name (%s) - new value (Enter to keep): ", existing.Name)
	if val := v.readLine(); val != "" {
		existing.Name = val
	}
	fmt.Printf("Email (%s) - new value (Enter to keep): ", existing.Email)
	if val := v.readLine(); val != "" {
		existing.Email = val
	}
	fmt.Printf("Phone (%s) - new value (Enter to keep): ", existing.Phone)
	if val := v.readLine(); val != "" {
		existing.Phone = val
	}
	fmt.Printf("City (%s) - new value (Enter to keep): ", existing.City)
	if val := v.readLine(); val != "" {
		existing.City = val
	}
	return existing
}

func (v *CustomerView) ShowCustomer(c models.Customer) {
	fmt.Printf("\nID: %d | Name: %s | Email: %s | Phone: %s | City: %s | Created: %s\n",
		c.ID, c.Name, c.Email, c.Phone, c.City, c.CreatedAt.Format("2006-01-02 15:04"))
}

func (v *CustomerView) ShowCustomers(customers []models.Customer, total, page, pageSize int) {
	if len(customers) == 0 {
		fmt.Println("\nNo customers found.")
		return
	}
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	fmt.Printf("\n--- Page %d of %d (Total: %d records) ---\n", page, totalPages, total)
	v.printTable(customers)
}

func (v *CustomerView) ShowSearchResults(customers []models.Customer) {
	if len(customers) == 0 {
		fmt.Println("\nNo matching customers found.")
		return
	}
	fmt.Printf("\nFound %d result(s):\n", len(customers))
	v.printTable(customers)
}

func (v *CustomerView) printTable(customers []models.Customer) {
	fmt.Printf("%-5s %-20s %-25s %-15s %-15s\n", "ID", "Name", "Email", "Phone", "City")
	fmt.Println(strings.Repeat("-", 85))
	for _, c := range customers {
		fmt.Printf("%-5d %-20s %-25s %-15s %-15s\n", c.ID, c.Name, c.Email, c.Phone, c.City)
	}
}

func (v *CustomerView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *CustomerView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
