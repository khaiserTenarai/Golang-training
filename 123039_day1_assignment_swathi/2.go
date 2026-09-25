
package main

import "fmt"

func main() {

	const applicationName = "Hello Cloud-Native Go"
	const applicationVersion = "1.0.0"
	const goVersion = "Go 1.25"
	const environment = "Development"

	fmt.Println("===== APPLICATION DETAILS =====")
	fmt.Println("Application Name:", applicationName)
	fmt.Println("Application Version:", applicationVersion)
	fmt.Println("Go Version:", goVersion)
	fmt.Println("Environment:", environment)
}


