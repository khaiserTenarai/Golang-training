package main
import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"money_transfer/controller"
	"money_transfer/database"
	"money_transfer/repository"
	"money_transfer/service"
	"money_transfer/view"
)

func main() {

	// Load environment variables.
	err := godotenv.Load("env/config.env")

	if err != nil {
		log.Fatal("Unable to load config.env")
	}

	// Connect database.
	db, err := database.Connect()

	if err != nil {
		log.Fatal(
			"Database connection failed:",
			err,
		)
	}

	defer db.Close()

	fmt.Println(
		"Database connected successfully.",
	)

	// Repository
	transferRepository :=
		repository.NewTransferRepository(db)

	// Service
	transferService :=
		service.NewTransferService(
			transferRepository,
		)

	// View
	transferView :=
		view.NewTransferView()

	// Controller
	transferController :=
		controller.NewTransferController(
			transferView,
			transferService,
		)

	// Start application
	transferController.Start()
}
