package main

import "fmt"

const CollegeName = "BTI"

func main() {
	var studentCount int = 25

	var isOpenforadmission = true

	department := "CS"

	var minPlacementPackage, maxPlacementPackage float64 = 300000, 1000000

	fmt.Println("College Name     :", CollegeName)
	fmt.Println("Number of students present in the college  :", studentCount)
	fmt.Println("Is taking admissions?    :", isOpenforadmission)
	fmt.Println("Department  :", department)
	fmt.Println("The minimum package placed in this college  :", minPlacementPackage)
	fmt.Println("he maximum package placed in this college   :", maxPlacementPackage)
}
