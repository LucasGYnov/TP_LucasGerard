package exercice1

func Ft_coin(coins []int, amount int) int {
	result := 0
	long := len(coins)
	rendu := make([]int, len(coins))

	if amount != 0 {
		for i := long - 1; i >= 0; i-- {
			for amount >= coins[i] {
				amount -= coins[i]
				rendu[i]++
				result = len(rendu)

			}
		}
		if amount != 0 {
			return -1
		}
	}
	return result
}
