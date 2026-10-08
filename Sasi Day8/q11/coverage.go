package q11

func DiscountedPrice(price float64, member bool) float64 {
	if price < 0 {
		return 0
	}
	if member {
		return price * .90
	}
	return price
}