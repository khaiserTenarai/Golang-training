package view
import "cms/model"

type CustomerView interface {

	ShowMenu() int

	ReadCustomer() model.Customer

	ReadCustomerForUpdate() model.Customer

	ReadID() int

	ReadSearchKeyword() string

	ReadPagination() (int, int)

	DisplayCustomer(customer model.Customer)

	DisplayCustomers(customers []model.Customer)
}
