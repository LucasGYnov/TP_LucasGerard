package main

import (
	exercice1 "exam/Exercice_1"
	exercice2 "exam/Exercice_2"
	exercice3 "exam/Exercice_3"
	exercice4 "exam/Exercice_4"
	exercice5 "exam/Exercice_5"
	exercice6 "exam/Exercice_6"
	"fmt"
)

// Test de la fonction Ft_coin
func testFtCoin() {
	fmt.Println("\nTest de la fonction Ft_coin :")
	fmt.Println(exercice1.Ft_coin([]int{1, 2, 5}, 11))
	fmt.Println(exercice1.Ft_coin([]int{2}, 3))
	fmt.Println(exercice1.Ft_coin([]int{1}, 0))
	fmt.Println("-------------------------------------")
}

// Test de la fonction Ft_missing
func testFtMissing() {
	fmt.Println("\nTest de la fonction Ft_missing :")
	fmt.Println(exercice2.Ft_missing([]int{3, 1, 2}))
	fmt.Println(exercice2.Ft_missing([]int{0, 1}))
	fmt.Println(exercice2.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println("-------------------------------------")
}

// Test de la fonction Ft_non_overlap
func testFtNonOverlap() {
	fmt.Println("\nTest de la fonction Ft_non_overlap :")
	fmt.Println(exercice3.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(exercice3.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))
	fmt.Println(exercice3.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println("-------------------------------------")
}

// Test de la fonction Ft_profit
func testFtProfit() {
	fmt.Println("\nTest de la fonction Ft_profit :")
	fmt.Println(exercice4.Ft_profit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(exercice4.Ft_profit([]int{7, 6, 4, 3, 1}))
	fmt.Println("-------------------------------------")
}

// Test de la fonction Ft_max_substring
func testFtMaxSubstring() {
	fmt.Println("\nTest de la fonction Ft_max_substring :")
	fmt.Println(exercice5.Ft_max_substring("abcabcbb"))
	fmt.Println(exercice5.Ft_max_substring("bbbbb"))
	fmt.Println("-------------------------------------")
}

// Test de la fonction Ft_min_window
func testFtMinWindow() {
	fmt.Println("\nTest de la fonction Ft_min_window :")
	fmt.Println(exercice6.Ft_min_window("ADOBECODEBANC", "ABC"))
	fmt.Println(exercice6.Ft_min_window("a", "aa"))
	fmt.Println("-------------------------------------")
}

func main() {
	var choice int
	for {
		fmt.Println("\n--- Menu de Test ---")
		fmt.Println("1. Tester Ft_coin")
		fmt.Println("2. Tester Ft_missing")
		fmt.Println("3. Tester Ft_non_overlap")
		fmt.Println("4. Tester Ft_profit")
		fmt.Println("5. Tester Ft_max_substring")
		fmt.Println("6. Tester Ft_min_window")
		fmt.Println("7. Tout tester")
		fmt.Println("0. Quitter")
		fmt.Print("\nFaites votre choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			testFtCoin()
		case 2:
			testFtMissing()
		case 3:
			testFtNonOverlap()
		case 4:
			testFtProfit()
		case 5:
			testFtMaxSubstring()
		case 6:
			testFtMinWindow()
		case 7:
			testFtCoin()
			testFtMissing()
			testFtNonOverlap()
			testFtProfit()
			testFtMaxSubstring()
			testFtMinWindow()
		case 0:
			fmt.Println("Merci, au revoir!")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
