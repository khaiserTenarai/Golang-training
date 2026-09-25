package main

import (
	"fmt"
	"os"
)

func main() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "MyApp"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Application:", appName)
	fmt.Println("Port:", port)
}


/*
Execution & Output :
------------------------

PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\10-environment-vars-demo> go mod init 10-environment-vars-demo
go: creating new go.mod: module 10-environment-vars-demo
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\10-environment-vars-demo> go run .
Application: MyApp
Port: 8080
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\10-environment-vars-demo> 
*/