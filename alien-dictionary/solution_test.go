package alien_dictionary

import "testing"

func TestIsAlienSorted(t *testing.T) {
	words := []string{"hello", "leetcode"}
	var order = "hlabcdefgijkmnopqrstuvwxyz"
	if !IsAlienSorted(words, order) {
		t.Errorf("wrong")
	}
}

func TestIsAlienSorted1(t *testing.T) {
	words := []string{"word", "world", "row"}
	var order = "worldabcefghijkmnpqstuvxyz"
	if IsAlienSorted(words, order) {
		t.Errorf("wrong")
	}
}

func TestIsAlienSorted2(t *testing.T) {
	words := []string{"apple", "app"}
	var order = "abcdefghijklmnopqrstuvwxyz"
	if IsAlienSorted(words, order) {
		t.Errorf("wrong")
	}
}

func TestIsAlienSorted3(t *testing.T) {
	words := []string{"kuvp", "q"}
	var order = "ngxlkthsjuoqcpavbfdermiywz"
	if !IsAlienSorted(words, order) {
		t.Errorf("wrong")
	}
}

