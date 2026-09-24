package main

import "fmt"

func main() {
    
    name := "लक्ष्मी" 

    byteCount := len(name) 
    
    characterList := []rune(name)
    characterCount := len(characterList)

    fmt.Println("Word:", name)
    fmt.Println("Bytes:", byteCount)
    fmt.Println("Characters:", characterCount)
}