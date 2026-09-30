package main
import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"bank_account/controller"
	"bank_account/database"
	"bank_account/repository"
	"bank_account/service"
	"bank_account/view"
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

	accountRepository :=
		repository.NewAccountRepository(db)

	accountService :=
		service.NewAccountService(
			accountRepository,
		)

	accountView :=
		view.NewAccountView()

	accountController :=
		controller.NewAccountController(
			accountView,
			accountService,
		)

	accountController.Start()
}
