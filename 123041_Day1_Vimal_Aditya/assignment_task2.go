package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("\n2. Hello Cloud-Native Go")
	AppName := "Cloud Native-Go"
	AppVersion := "v1.5"
	EnvName := "Production"
	goVersion := runtime.Version()

	fmt.Println("========================================")
	fmt.Println("Application Name : ", AppName)
	fmt.Println("App Version      : ", AppVersion)
	fmt.Println("Go Version       : ", goVersion)
	fmt.Println("Environment      : ", EnvName)
	fmt.Println("========================================")

}