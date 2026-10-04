package main

import (
	"context"
	"fmt"

	"ecommerce-management/config"
	"ecommerce-management/controller"
	"ecommerce-management/database"
	"ecommerce-management/repository"
	"ecommerce-management/service"
	"ecommerce-management/view"
)

func main() {

	// Load .env configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db, err := database.ConnectDB(cfg)

	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	defer db.Close(context.Background())

	fmt.Println("Database connected successfully")

	// Create repository
	repo := repository.NewRepository(db)

	// Create service
	svc := service.NewService(repo)

	// Create controller
	ctrl := controller.NewController(svc)

	// Menu loop
	for {

		view.ShowMenu()

		choice := view.GetChoice()

		switch choice {

		case 1:
			ctrl.AddProduct()

		case 2:
			ctrl.DisplayProducts()

		case 3:
			ctrl.UpdateProduct()

		case 4:
			ctrl.DeleteProduct()

		case 5:
			ctrl.IncreaseStock()

		case 6:
			ctrl.DecreaseStock()

		case 7:
			ctrl.LowStockProducts()

		case 8:
			fmt.Println("Exiting application...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
