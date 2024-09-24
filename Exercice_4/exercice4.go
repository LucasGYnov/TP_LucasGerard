package Exercice_4

func Ft_profit(prices []int) int {
	// Si le tableau de prix est vide, il n'y a aucun bénéfice
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0] // Initialiser le prix minimum
	maxProfit := 0

	// Parcourir les prix à partir du deuxième élément
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
