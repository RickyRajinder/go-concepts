package kClosest

import "testing"

func TestKClosest1(t *testing.T) {
	res := kClosest([][]int{{1,3}, {-2, 2}, {-2, 2}}, 1)
	if len(res) != 1 {
		t.Errorf("wrong")
	}
}

func TestKClosest(t *testing.T) {
	res := kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)
	if len(res) != 2 {
		t.Errorf("wrong")
	}
}
