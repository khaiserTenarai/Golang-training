package controller

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "customer-management/model"
    "customer-management/service"
)

type CustomerController struct {
    customerService service.CustomerService
    reader          *bufio.Reader
}

func NewCustomerController(customerService service.CustomerService) *CustomerController {
    return &CustomerController{
        customerService: customerService,
        reader:          bufio.NewReader(os.Stdin),
    }
}

func (c *CustomerController) Start() {
    for {
        fmt.Println()
        fmt.Println("========== CUSTOMER MANAGEMENT ==========")
        fmt.Println("1. Add Customer")
        fmt.Println("2. Get Customer by ID")
        fmt.Println("3. View Customers")
        fmt.Println("4. Update Customer")
        fmt.Println("5. Delete Customer")
        fmt.Println("6. Search Customer")
        fmt.Println("7. Exit")
        fmt.Println("=========================================")

        choice := c.readInt("Enter choice: ")

        switch choice {
        case 1:
            c.addCustomer()
        case 2:
            c.getCustomer()
        case 3:
            c.viewCustomers()
        case 4:
            c.updateCustomer()
        case 5:
            c.deleteCustomer()
        case 6:
            c.searchCustomers()
        case 7:
            fmt.Println("Thank you for using Customer Management System.")
            return
        default:
            fmt.Println("Invalid choice.")
        }
    }
}

func (c *CustomerController) addCustomer() {
    customer := model.Customer{
        Name:  c.readString("Enter name: "),
        Email: c.readString("Enter email: "),
        Phone: c.readString("Enter phone: "),
        City:  c.readString("Enter city: "),
    }

    if err := c.customerService.AddCustomer(customer); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Customer added successfully.")
}

func (c *CustomerController) getCustomer() {
    id := c.readInt("Enter customer ID: ")

    customer, err := c.customerService.GetCustomerByID(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    printCustomer(customer)
}

func (c *CustomerController) viewCustomers() {
    page := c.readInt("Enter page number: ")
    pageSize := c.readInt("Enter page size: ")

    if page <= 0 {
        page = 1
    }
    if pageSize <= 0 {
        pageSize = 5
    }

    offset := (page - 1) * pageSize

    customers, err := c.customerService.GetAllCustomers(pageSize, offset)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(customers) == 0 {
        fmt.Println("No customers found on this page.")
        return
    }

    fmt.Printf("
Page %d
", page)
    printCustomers(customers)
}

func (c *CustomerController) updateCustomer() {
    id := c.readInt("Enter customer ID: ")

    customer, err := c.customerService.GetCustomerByID(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Press Enter to keep the existing value.")

    name := c.readOptionalString("Name [" + customer.Name + "]: ")
    email := c.readOptionalString("Email [" + customer.Email + "]: ")
    phone := c.readOptionalString("Phone [" + customer.Phone + "]: ")
    city := c.readOptionalString("City [" + customer.City + "]: ")

    if name != "" {
        customer.Name = name
    }
    if email != "" {
        customer.Email = email
    }
    if phone != "" {
        customer.Phone = phone
    }
    if city != "" {
        customer.City = city
    }

    if err := c.customerService.UpdateCustomer(customer); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Customer updated successfully.")
}

func (c *CustomerController) deleteCustomer() {
    id := c.readInt("Enter customer ID: ")

    confirm := strings.ToLower(c.readString("Are you sure? (yes/no): "))
    if confirm != "yes" && confirm != "y" {
        fmt.Println("Delete cancelled.")
        return
    }

    if err := c.customerService.DeleteCustomer(id); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Customer deleted successfully.")
}

func (c *CustomerController) searchCustomers() {
    keyword := c.readString("Enter search keyword: ")
    page := c.readInt("Enter page number: ")
    pageSize := c.readInt("Enter page size: ")

    if page <= 0 {
        page = 1
    }
    if pageSize <= 0 {
        pageSize = 5
    }

    offset := (page - 1) * pageSize

    customers, err := c.customerService.SearchCustomers(keyword, pageSize, offset)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(customers) == 0 {
        fmt.Println("No matching customers found.")
        return
    }

    fmt.Printf("
Search results - Page %d
", page)
    printCustomers(customers)
}

func (c *CustomerController) readString(prompt string) string {
    fmt.Print(prompt)
    value, _ := c.reader.ReadString('
')
    return strings.TrimSpace(value)
}

func (c *CustomerController) readOptionalString(prompt string) string {
    return c.readString(prompt)
}

func (c *CustomerController) readInt(prompt string) int {
    for {
        value := c.readString(prompt)
        number, err := strconv.Atoi(value)
        if err == nil {
            return number
        }
        fmt.Println("Please enter a valid number.")
    }
}

func printCustomer(customer model.Customer) {
    fmt.Println("
---------- CUSTOMER ----------")
    fmt.Println("ID:", customer.ID)
    fmt.Println("Name:", customer.Name)
    fmt.Println("Email:", customer.Email)
    fmt.Println("Phone:", customer.Phone)
    fmt.Println("City:", customer.City)
    fmt.Println("Created At:", customer.CreatedAt.Format("2006-01-02 15:04:05"))
    fmt.Println("------------------------------")
}

func printCustomers(customers []model.Customer) {
    fmt.Println("
ID | Name | Email | Phone | City")
    fmt.Println("-------------------------------------------------------------")

    for _, customer := range customers {
        fmt.Printf("%d | %s | %s | %s | %s
",
            customer.ID,
            customer.Name,
            customer.Email,
            customer.Phone,
            customer.City,
        )
    }
}
