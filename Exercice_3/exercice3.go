package Exercice_3

import "sort"

func Ft_non_overlap(intervals [][]int) int {
	// Si le tableau d'intervalles est vide, il n'y a rien à retirer.
	if len(intervals) == 0 {
		return 0
	}

	// Trier les intervalles
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0             // Initalisation d'un compteur pour le nombre d'intervalles à retirer
	end := intervals[0][1] // Initialiser 'end' avec la fin du premier intervalle

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}
