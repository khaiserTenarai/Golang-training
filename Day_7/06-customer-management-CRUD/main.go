package main
import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"cms/controller"
	"cms/database"
	"cms/repository"
	"cms/service"
	"cms/view"
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

	customerController.Start()
}
