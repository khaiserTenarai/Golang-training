
package main

import "fmt"

func main() {

	name := "Swathi"
	indianName := "ಕಾವ್ಯ"

	fmt.Println("===== BYTES VS RUNES =====")

	fmt.Println("Name:", name)
	fmt.Println("Indian Name:", indianName)

	fmt.Println("\n===== BYTES =====")

	for i := 0; i < len(indianName); i++ {
		fmt.Println(indianName[i])
	}

	fmt.Println("\n===== RUNES =====")

	for _, character := range indianName {
		fmt.Println(character)
	}

	fmt.Println("\n===== LENGTH =====")
	fmt.Println("Byte length:", len(indianName))
	fmt.Println("Rune count:", len([]rune(indianName)))
}

