package main

import "fmt"

func main() {

	fmt.Println("Hello...")
}

/*
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> go mod init 13-simple-automation
go: creating new go.mod: module 13-simple-automation
go: to add module requirements and sums:
        go mod tidy
*/


/*
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> .\build.ps1 build                                         
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> .\build.ps1 run                                           
Hello...
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> .\build.ps1 vet  
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> .\build.ps1 test
?       13-simple-automation    [no test files]
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> .\build.ps1 fmt 
main.go
PS C:\Training\Go Lang\Day_1\124772_Day1_Assignment_Reddem_Ganesh_Reddy\13-simple-automation> 
*/