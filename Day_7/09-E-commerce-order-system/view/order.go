package view
import (
	"fmt"

	"ecommerce/model"
)

type OrderView interface {
	ShowMenu() int
	ReadOrder() model.OrderRequest
	ReadID() int
	DisplayOrder(details []model.OrderDetail)
	DisplayOrders(details []model.OrderDetail)
}

type OrderViewImpl struct {
}

func NewOrderView() OrderView {
	return &OrderViewImpl{}
}

func (v *OrderViewImpl) ShowMenu() int {

	fmt.Println("\n========== Order Management ==========")
	fmt.Println("1. Create Order")
	fmt.Println("2. Find Order")
	fmt.Println("3. Find All Orders")
	fmt.Println("4. Back")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *OrderViewImpl) ReadOrder() model.OrderRequest {

	var order model.OrderRequest

	fmt.Println("\n---------- Create Order ----------")

	fmt.Print("Enter Customer ID: ")
	fmt.Scan(&order.CustomerID)

	var count int

	fmt.Print("Enter number of products: ")
	fmt.Scan(&count)

	order.Items = make(
		[]model.OrderItemRequest,
		0,
		count,
	)

	for i := 0; i < count; i++ {

		var item model.OrderItemRequest

		fmt.Println("\nOrder Item", i+1)

		fmt.Print("Enter Product ID: ")
		fmt.Scan(&item.ProductID)

		fmt.Print("Enter Quantity: ")
		fmt.Scan(&item.Quantity)

		order.Items = append(
			order.Items,
			item,
		)
	}

	return order
}

func (v *OrderViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Order ID: ")
	fmt.Scan(&id)

	return id
}

func (v *OrderViewImpl) DisplayOrder(
	details []model.OrderDetail,
) {

	if len(details) == 0 {
		fmt.Println("No order found.")
		return
	}

	first := details[0]

	fmt.Println("\n==========================================")
	fmt.Println("              ORDER DETAILS")
	fmt.Println("==========================================")

	fmt.Println("Order ID       :", first.OrderID)
	fmt.Println("Order Date     :", first.OrderDate)
	fmt.Println("Customer ID    :", first.CustomerID)
	fmt.Println("Customer Name  :", first.CustomerName)
	fmt.Println("Customer Email :", first.CustomerEmail)
	fmt.Println("Customer Phone :", first.CustomerPhone)

	fmt.Println("------------------------------------------")

	var grandTotal float64

	for _, detail := range details {

		fmt.Println("Order Item ID :", detail.OrderItemID)
		fmt.Println("Product ID    :", detail.ProductID)
		fmt.Println("Product Name  :", detail.ProductName)
		fmt.Println("Price         :", detail.Price)
		fmt.Println("Quantity      :", detail.Quantity)
		fmt.Println("Item Total    :", detail.ItemTotal)

		fmt.Println("------------------------------------------")

		grandTotal += detail.ItemTotal
	}

	fmt.Println("Grand Total   :", grandTotal)
	fmt.Println("==========================================")
}

func (v *OrderViewImpl) DisplayOrders(
	details []model.OrderDetail,
) {

	if len(details) == 0 {
		fmt.Println("No orders found.")
		return
	}

	fmt.Println("\n================ ALL ORDERS ================")

	currentOrderID := -1
	var grandTotal float64

	for _, detail := range details {

		if currentOrderID != detail.OrderID {

			if currentOrderID != -1 {
				fmt.Println("Grand Total:", grandTotal)
				fmt.Println("--------------------------------------------")
			}

			currentOrderID = detail.OrderID
			grandTotal = 0

			fmt.Println("\nOrder ID      :", detail.OrderID)
			fmt.Println("Order Date    :", detail.OrderDate)
			fmt.Println("Customer Name :", detail.CustomerName)
			fmt.Println("Customer Email:", detail.CustomerEmail)
		}

		fmt.Println(
			"  Product:",
			detail.ProductName,
			"| Price:",
			detail.Price,
			"| Quantity:",
			detail.Quantity,
			"| Total:",
			detail.ItemTotal,
		)

		grandTotal += detail.ItemTotal
	}

	fmt.Println("Grand Total:", grandTotal)
}
