package rm_invalid_paren

import "testing"

func TestRmParen(t *testing.T) {
	res := RemoveInvalidParentheses("()())()")
	if !testArrayEquals(res, []string{"(())()", "()()()"}) {
		t.Errorf("wrong")
	}
}

func TestRmParen1(t *testing.T) {
	res := RemoveInvalidParentheses("(a)())()")
	if !testArrayEquals(res, []string{"(a())()", "(a)()()"}) {
		t.Errorf("wrong")
	}
}

func TestRmParen2(t *testing.T) {
	res := RemoveInvalidParentheses(")(")
	if !testArrayEquals(res, []string{""}) {
		t.Errorf("wrong")
	}
}



func testArrayEquals(a, b []string) bool {
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

