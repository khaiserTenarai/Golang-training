package main

import "fmt"

func main() {
    var choice int

    fmt.Println("1. Say Hello")
    fmt.Println("2. Say Goodbye")
    fmt.Print("Choose 1 or 2: ")

    fmt.Scan(&choice)

    switch choice {
    case 1:
        fmt.Println("Hello")
    case 2:
        fmt.Println("Goodbye")
    default:
        fmt.Println("Wrong choice")
    }
}