package exercice6

func Ft_min_window(source string, target string) string {
	// Vérification des erreurs, si la chaîne 'target' est vide
	if len(target) == 0 {
		return ""
	}

	// Dictionnaire pour compter les occurrences de chaque caractère dans 'target'
	countTarget := make(map[byte]int)
	for i := range target {
		countTarget[target[i]]++
	}

	presentCount, requiredCount := 0, len(countTarget) // var 'presentCount' : nombre de caractères de 'target' dans la fenêtre, var 'requiredCount' : nombre total de caractères nécessaires
	windowCount := make(map[byte]int)                  // Création d'un dictionnaire pour suivre les caractères dans la fenêtre
	bestSubstring, minLength := "", len(source)+1

	leftIndex := 0 // Indice de gauche

	for rightIndex := 0; rightIndex < len(source); rightIndex++ {
		currentChar := source[rightIndex] // Caractère courant
		windowCount[currentChar]++        // Ajouter le caractère courant

		if count, ok := countTarget[currentChar]; ok && windowCount[currentChar] == count {
			presentCount++ // Incrémenter le nombre de caractères  requis
		}

		for presentCount == requiredCount {
			if (rightIndex - leftIndex + 1) < minLength {
				bestSubstring = source[leftIndex : rightIndex+1] // Màj la meilleure sous-chaîne
				minLength = rightIndex - leftIndex + 1           // Màj la longueur de la meilleure sous-chaîne
			}

			// Réduire la taille de la fenêtre en retirant le caractère à gauche.
			windowCount[source[leftIndex]]--
			if count, ok := countTarget[source[leftIndex]]; ok && windowCount[source[leftIndex]] < count {
				presentCount-- // Décrémenter le nombre de caractères requis
			}
			leftIndex++ // Avancer l'indice de gauche
		}
	}

	// Si aucune sous-chaîne n'a été trouvée, retourner une chaîne vide
	if minLength == len(source)+1 {
		return ""
	}
	return bestSubstring
}
