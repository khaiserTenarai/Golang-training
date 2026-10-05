// 11. Measure test coverage.

package main

func Grade(marks int) string {
	switch {
	case marks >= 90:
		return "A"
	case marks >= 75:
		return "B"
	case marks >= 50:
		return "C"
	default:
		return "F"
	}
}
