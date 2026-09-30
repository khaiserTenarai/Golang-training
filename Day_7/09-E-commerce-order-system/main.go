package main
import (
	"fmt"
	"os"

	"ecommerce/controller"
	"ecommerce/database"
	"ecommerce/repository"
	"ecommerce/service"
	"ecommerce/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println(
			"Database connection failed:",
			err,
		)

		os.Exit(1)
	}

	defer db.Close()

	// -------------------------
	// Customer
	// -------------------------

	customerRepository :=
		repository.NewCustomerRepository(db)

	customerService :=
		service.NewCustomerService(
			customerRepository,
		)

	customerView :=
		view.NewCustomerView()

	customerController :=
		controller.NewCustomerController(
			customerView,
			customerService,
		)

	// -------------------------
	// Product
	// -------------------------

	productRepository :=
		repository.NewProductRepository(db)

	productService :=
		service.NewProductService(
			productRepository,
		)

	productView :=
		view.NewProductView()

	productController :=
		controller.NewProductController(
			productView,
			productService,
		)

	// -------------------------
	// Order
	// -------------------------

	orderRepository :=
		repository.NewOrderRepository(db)

	orderService :=
		service.NewOrderService(
			orderRepository,
		)

	orderView :=
		view.NewOrderView()

	orderController :=
		controller.NewOrderController(
			orderView,
			orderService,
		)

	// -------------------------
	// Main Menu
	// -------------------------

	for {

		fmt.Println("\n==========================================")
		fmt.Println("       E-COMMERCE ORDER SYSTEM")
		fmt.Println("==========================================")
		fmt.Println("1. Customer Management")
		fmt.Println("2. Product Management")
		fmt.Println("3. Order Management")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			customerController.Start()

		case 2:

			productController.Start()

		case 3:

			orderController.Start()

		case 4:

			fmt.Println("Thank you..")
			return

		default:

			fmt.Println("Invalid choice.")
		}
	}
}
