package slidingwindow

func BestTimeForTransaction(prices []int) int {
	minPricePtr := 0
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[minPricePtr] >= prices[i] {
			minPricePtr = i
		} else {
			maxProfit = max(maxProfit, prices[i]-prices[minPricePtr])
		}
	}

	return maxProfit
}
