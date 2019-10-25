package merge

import "testing"

func TestMergeLists(t *testing.T) {
	arr1 := []int{1, 2, 3, 5}
	arr2 := []int{2, 2, 4, 5}
	merged := mergeLists(arr1, arr2)
	expected := []int{1, 2, 2, 2, 3, 4, 5, 5}
	if !testArrayEquals(merged, expected) {
		t.Errorf("Arrays not the same")
	}
}

func testArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, j := range a {
		if j != b[i] {
			return false
		}
	}
	return true
}
