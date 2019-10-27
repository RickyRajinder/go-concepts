package bst

import "testing"

func TestContainsInserted(t *testing.T) {
	bt := createTestBinaryTree()
	if !bt.Contains(6) {
		t.Errorf("fail")
	}
	if bt.Contains(99) {
		t.Errorf("fail")
	}
	if !bt.Contains(7) {
		t.Errorf("fail")
	}
	if !bt.Contains(3) {
		t.Errorf("fail")
	}
}
