package exercice5

func Ft_max_substring(s string) int {
	m := make(map[byte]int)
	maxLen := 0
	start := 0

	// Parcourir chaque caractère
	for i := 0; i < len(s); i++ {
		if pos, ok := m[s[i]]; ok && pos >= start {
			start = pos + 1
		}
		// Màj de la position du caractère
		m[s[i]] = i
		// Calculer la longueur de la chaîne et màj la longueur maximale
		maxLen = max(maxLen, i-start+1)
	}

	return maxLen
}

// Fonction max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
