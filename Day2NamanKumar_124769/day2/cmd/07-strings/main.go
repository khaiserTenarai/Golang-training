// Task 7: String-processing application.
// Counts characters, words, vowels and digits (plus a few extras).
//
// Usage:
//
//	go run ./cmd/07-strings                    (type text, empty line to finish)
//	go run ./cmd/07-strings "some text here"   (analyse the argument)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"day2/internal/textstats"
)

func main() {
	var text string
	if len(os.Args) > 1 {
		text = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Println("Enter text (press Enter on an empty line to finish):")
		var lines []string
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() && sc.Text() != "" {
			lines = append(lines, sc.Text())
		}
		text = strings.Join(lines, "\n")
	}
	if strings.TrimSpace(text) == "" {
		text = "Go was created at Google in 2009. Version 1.22 released in 2024!"
		fmt.Println("(no input, using sample text)")
	}

	s := textstats.Analyze(text)
	fmt.Printf("\nText: %q\n\n", text)
	fmt.Println("========== TEXT STATISTICS ==========")
	fmt.Printf("%-24s %d\n", "Characters (with spaces)", s.Characters)
	fmt.Printf("%-24s %d\n", "Characters (no spaces)", s.NonSpace)
	fmt.Printf("%-24s %d\n", "Words", s.Words)
	fmt.Printf("%-24s %d\n", "Vowels", s.Vowels)
	fmt.Printf("%-24s %d\n", "Digits", s.Digits)
	fmt.Println("-------------------------------------")
	fmt.Printf("%-24s %d\n", "Consonants", s.Consonants)
	fmt.Printf("%-24s %d\n", "Spaces", s.Spaces)
	fmt.Printf("%-24s %d\n", "Punctuation", s.Punctuation)
	fmt.Printf("%-24s %d\n", "Lines", s.Lines)
	fmt.Printf("%-24s %d\n", "Bytes (len)", s.Bytes)

	fmt.Println("\n=== Useful strings functions ===")
	fmt.Println("Upper   :", strings.ToUpper(text))
	fmt.Println("Words   :", strings.Fields(text))
	fmt.Println("Has 'Go':", strings.Contains(text, "Go"))
	fmt.Println("Reversed:", reverse(text))
}

// reverse reverses a string rune by rune, so non-English text stays valid.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
