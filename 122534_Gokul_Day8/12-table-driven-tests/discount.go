// 12. Create table-driven tests.

package main

func DiscountPercent(purchaseAmount float64) float64 {
	switch {
	case purchaseAmount >= 10000:
		return 20
	case purchaseAmount >= 5000:
		return 10
	case purchaseAmount >= 1000:
		return 5
	default:
		return 0
	}
}
