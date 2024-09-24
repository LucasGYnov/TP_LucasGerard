package main

import (
	exercice1 "exam/Exercice_1"
	"fmt"
)

func main() {
	/* Exercice 1 */
	fmt.Print("\n-----------------------\n" + "Exercice 1 : Test de la fonction Ft_coin\n")
	fmt.Println(exercice1.Ft_coin([]int{1, 2, 5}, 11))
	fmt.Println(exercice1.Ft_coin([]int{2}, 3))
	fmt.Println(exercice1.Ft_coin([]int{1}, 0))

	/* Exercice 2 */
	fmt.Print("\n-----------------------\n" + "Exercice 2 : Test de la fonction Ft_missing\n")

}
