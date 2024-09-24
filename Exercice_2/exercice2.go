package exercice2

import "sort"

func Ft_missing(nums []int) int {
	long := len(nums)
	result := 0

	// Tri du tableau nums
	sort.Ints(nums)

	// Vérification du nombre manquant
	for i := 0; i < long; i++ {
		if nums[i] != i {
			result = i
			break
		}
	}

	// Si aucun nombre manquant n'a été trouvé, le resultat est assigné à long
	if result == 0 && nums[long-1] == long-1 {
		result = long
	}

	return result
}
