package subarray_sum_equals_k


import "testing"

func TestRmParen(t *testing.T) {
	res := subarraySum([]int{1,1,1}, 2)
	if res != 2 {
		t.Errorf("wrong")
	}
}
