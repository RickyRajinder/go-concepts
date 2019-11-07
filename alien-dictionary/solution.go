package alien_dictionary

import (
	"math"
)

//Time O(chars in words), Space(1)
func IsAlienSorted(words []string, order string) bool {
	var count = make([]int, 26)

	for _, j := range order {
		count[j - 'a']++
	}

	for k := 0; k < len(words)-1; k++ {
		word1 := words[k]
		word2 := words[k+1]

		for i := 0; i < int(math.Min(float64(len(word1)), float64(len(word2)))); i++ {
			char1 := word1[i]
			char2 := word2[i]
			if char1 != char2 {
				if count[char1 - 'a'] > count[char2 - 'a'] {
					return false
				}
				continue
			}
		}

		if len(word1) > len(word2) {
			return false
		}
	}
	return true
}
