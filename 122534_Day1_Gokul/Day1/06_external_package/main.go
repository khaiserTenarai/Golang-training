
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// generate a random unique id
	id := uuid.New()
	
	fmt.Println("new record created successfullcd")
	fmt.Println("unique id:", id.String())
}
