package alien_dictionary

//Time O(chars in words), Space(1)
func IsAlienSorted(words []string, order string) bool {
	var count = make([]int, 26)

	for j := 0; j < len(order); j++ {
		count[order[j] - 'a'] = j
	}

	loop1: for k := 0; k < len(words)-1; k++ {
		word1 := words[k]
		word2 := words[k+1]

		for i := 0; i < Min(len(word1), len(word2)); i++ {
			char1 := word1[i]
			char2 := word2[i]
			if char1 != char2 {
				if count[char1 - 'a'] > count[char2 - 'a'] {
					return false
				}
				continue loop1
			}
		}

		if len(word1) > len(word2) {
			return false
		}
	}
	return true
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
