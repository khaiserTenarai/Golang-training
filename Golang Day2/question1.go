package main

import "fmt"

func main() {

	const appName ="Go App"
	const releaseYear=2026
	var username string ="Lakshmi"
	var currentScore int=24

	isLoggedIn:=true

	fmt.Println("App Name:", appName)
	fmt.Println("Year:", releaseYear)
	fmt.Println("User:", username)
	fmt.Println("Logged In:", isLoggedIn)
	fmt.Println("Score:", currentScore)
	 
	currentScore=45
	fmt.Println("New Score:", currentScore)
}