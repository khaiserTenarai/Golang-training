package main
import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"product_inventory/controller"
	"product_inventory/database"
	"product_inventory/repository"
	"product_inventory/service"
	"product_inventory/view"
)

func main() {

	err := godotenv.Load("env/config.env")

	if err != nil {
		log.Fatal("Unable to load config.env")
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer db.Close()

	fmt.Println("Database connected successfully.")

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

	productController.Start()
}
