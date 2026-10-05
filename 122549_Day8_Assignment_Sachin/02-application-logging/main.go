// 2. Implement application logging.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("Application started")

	reader := bufio.NewReader(os.Stdin)

	log.Println("Waiting for employee name input")
	fmt.Print("Enter employee name: ")
	nameText, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameText)

	if name == "" {
		log.Println("No name entered, aborting")
		return
	}

	log.Printf("Employee added: %s\n", name)
	log.Println("Application finished")
}
