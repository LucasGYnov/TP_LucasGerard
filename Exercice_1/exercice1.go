package exercice1

import "sort"

func Ft_coin(coins []int, amount int) int {
	/***  VERFICATION DES SORTIES/ERREURS POSSIBLES  ***/

	// Si amount est  égal à 0, il n'y a pas besoin de pièces
	if amount == 0 {
		return 0
	}
	// Si le tableau de pièces est vide, il est impossible d'atteindre le montant
	if len(coins) == 0 {
		return -1
	}

	// Trier le tableau de pièces de façon décroissante, pour éviter les problèmes provenant du tableau
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	result := 0

	for _, coin := range coins {
		for amount >= coin {
			amount -= coin
			result++
		}
	}
	if amount != 0 {
		return -1
	}
	return result
}
