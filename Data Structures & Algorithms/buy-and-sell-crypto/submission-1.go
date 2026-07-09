func maxProfit(prices []int) int {
	maxProfit := 0
	minBuy := prices[0]
	for _, sell := range prices{
		if sell - minBuy > maxProfit{
			maxProfit = sell - minBuy
		}
		if sell < minBuy{
			minBuy = sell
		}
	}
	return maxProfit
}
