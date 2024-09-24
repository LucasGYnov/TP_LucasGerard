package Exercice_4

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		currentProfit := price - minPrice

		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}

		if price < minPrice {
			minPrice = price
		}
	}

	return maxProfit
}
