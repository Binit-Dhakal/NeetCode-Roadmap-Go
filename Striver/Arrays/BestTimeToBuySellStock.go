package arrays

func BestTime(prices []int) int {
	maxPrice := 0
	currPrice := prices[0]
	for _, price := range prices[1:] {
		if price < currPrice {
			currPrice = price
			continue
		}

		maxPrice = max(maxPrice, price-currPrice)
	}

	return maxPrice
}
