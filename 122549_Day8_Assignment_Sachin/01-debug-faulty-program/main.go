// 1. Debug a faulty Go program.
//
// The ORIGINAL, buggy version of this program looked like this:
//
//   func highestSalary(salaries []float64) float64 {
//       highest := salaries[0]
//       for i := 0; i <= len(salaries); i++ {   // BUG: should be i < len(salaries)
//           if salaries[i] > highest {
//               highest = salaries[i]
//           }
//       }
//       return highest
//   }
//
// Running it panics with "index out of range" once i reaches len(salaries),
// because slice indexes only go up to len(salaries)-1. The fix is just
// changing "<=" to "<" in the loop condition. The corrected, working
// version is below.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func highestSalary(salaries []float64) float64 {
	highest := salaries[0]
	for i := 0; i < len(salaries); i++ {
		if salaries[i] > highest {
			highest = salaries[i]
		}
	}
	return highest
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many employees? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive number.")
		return
	}

	salaries := make([]float64, 0, count)
	for i := 1; i <= count; i++ {
		fmt.Printf("Enter salary for employee %d: ", i)
		salText, _ := reader.ReadString('\n')
		salary, err := strconv.ParseFloat(strings.TrimSpace(salText), 64)
		if err != nil {
			fmt.Println("That's not a valid number, skipping.")
			continue
		}
		salaries = append(salaries, salary)
	}

	if len(salaries) == 0 {
		fmt.Println("No valid salaries entered.")
		return
	}

	fmt.Println("Highest salary:", highestSalary(salaries))
}
