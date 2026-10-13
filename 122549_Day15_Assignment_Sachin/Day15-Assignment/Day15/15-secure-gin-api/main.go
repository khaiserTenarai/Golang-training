package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(question string) string {
		fmt.Print(question)
		text, _ := reader.ReadString('\n')
		return strings.TrimSpace(text)
	}

	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	secret := ask("Enter a secret key used to sign tokens: ")
	adminName := ask("Create an admin username: ")
	adminPass := ask("Create an admin password: ")
	userName := ask("Create a normal user username: ")
	userPass := ask("Create a normal user password: ")

	if secret == "" || adminName == "" || adminPass == "" || userName == "" || userPass == "" {
		fmt.Println("Every value is needed.")
		return
	}
	if adminName == userName {
		fmt.Println("The two usernames must be different.")
		return
	}

	auth := NewAuth(secret)
	if err := auth.AddUser(adminName, adminPass, "admin"); err != nil {
		fmt.Println("Could not create the admin:", err)
		return
	}
	if err := auth.AddUser(userName, userPass, "user"); err != nil {
		fmt.Println("Could not create the user:", err)
		return
	}

	router := NewRouter(auth, NewStore())
	fmt.Println("Server running on port", port)
	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
