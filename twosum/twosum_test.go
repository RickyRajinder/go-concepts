package main

import "testing"

func Test(t *testing.T){
	res := twosum([]int{1, 2, 3, 4}, 3)
	if len(res) != 2 {
		t.Errorf("Incorrect length")
	}
	sol := []int{0, 1}
	for i ,v := range sol {
		if v != res[i]{
			t.Errorf("Incorrect indices")
		}
	}
}
