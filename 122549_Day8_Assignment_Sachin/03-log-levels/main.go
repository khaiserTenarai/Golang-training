// 3. Implement different log levels.
//
// The standard "log" package doesn't have built-in levels like
// DEBUG/INFO/WARN/ERROR, so this just prefixes each message by hand to
// tell them apart. Bigger projects usually use a proper logging library
// for this, but the idea is the same one shown here.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func logDebug(msg string) { log.Println("[DEBUG]", msg) }
func logInfo(msg string)  { log.Println("[INFO]", msg) }
func logWarn(msg string)  { log.Println("[WARN]", msg) }
func logError(msg string) { log.Println("[ERROR]", msg) }

func main() {
	reader := bufio.NewReader(os.Stdin)

	logInfo("Application started")

	fmt.Print("Enter employee age: ")
	ageText, _ := reader.ReadString('\n')
	age, err := strconv.Atoi(strings.TrimSpace(ageText))
	if err != nil {
		logError("Age entered is not a valid number")
		return
	}
	logDebug(fmt.Sprintf("Raw age value received: %d", age))

	switch {
	case age < 0:
		logError("Age cannot be negative")
	case age < 18:
		logWarn("Employee is under 18, needs special approval")
	default:
		logInfo("Age is valid")
	}
}
