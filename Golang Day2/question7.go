package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Print("Enter a sentence: ")
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()

    charCount := len(text)
    wordCount := len(strings.Fields(text)) 
    
    vowelCount := 0
    digitCount := 0

    text = strings.ToLower(text)

    for _, letter := range text {
        
        if letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' {
            vowelCount++
        }
        
        if letter >= '0' && letter <= '9' {
            digitCount++
        }
    }

    fmt.Println("Characters:", charCount)
    fmt.Println("Words     :", wordCount)
    fmt.Println("Vowels    :", vowelCount)
    fmt.Println("Digits    :", digitCount)
}