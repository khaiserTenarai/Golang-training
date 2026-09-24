package main

import(
	"fmt"
	"os"
	"runtime"
)

func main(){

	appName:="Cloud-Native App"
	appVersion:="v1.0.0"

	envName:=os.Getenv("APP_ENV")

	if envName==""{
	envName="development"
}
    goVersion:=runtime.Version()

	fmt.Println("Application Name    :", appName)
	fmt.Println("Application Version :", appVersion)
	fmt.Println("Go Version          :", goVersion)
	fmt.Println("Environment Name    :", envName)
}