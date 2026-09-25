package main

import (
	"fmt"
	"runtime"
)

type AppConfig struct {
	Name    string
	Version string
}

func ShowDetails(config AppConfig) {
	fmt.Println("App:", config.Name)
	fmt.Println("Ver:", config.Version)
	fmt.Println("Go :", runtime.Version())
}

func main() {
	myApp := AppConfig{Name: "VS Code", Version: "1.0"}
	ShowDetails(myApp)
}
