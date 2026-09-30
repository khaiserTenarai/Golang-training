package main
import (
	"fmt"
	"os"

	"order_inventory/controller"
	"order_inventory/database"
	"order_inventory/repository"
	"order_inventory/service"
	"order_inventory/view"
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

	// --------------------------------
	// PRODUCT
	// --------------------------------

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

	// --------------------------------
	// ORDER
	// --------------------------------

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

	// --------------------------------
	// MAIN MENU
	// --------------------------------

	for {

		fmt.Println("\n==========================================")
		fmt.Println("       ORDER & INVENTORY SYSTEM")
		fmt.Println("==========================================")
		fmt.Println("1. Product Management")
		fmt.Println("2. Order Management")
		fmt.Println("3. Exit")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			productController.Start()

		case 2:

			orderController.Start()

		case 3:

			fmt.Println("Thank you..")
			return

		default:

			fmt.Println("Invalid choice.")
		}
	}
}
