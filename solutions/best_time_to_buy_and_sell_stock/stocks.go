package besttimetobuyandsellstock

func maxProfit(prices []int) (maxProfit int) {
	lowest := prices[0]
	maxProfit = 0

	for _, price := range prices[1:] {
		if price < lowest {
			lowest = price
		} else if maxProfit < price-lowest {
			maxProfit = price - lowest
		}
	}

	return
}
