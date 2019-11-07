package rm_invalid_paren

import "strings"

type Solution struct {
	set map[string]bool
}

//O(2^N) time, O(N) space
func RemoveInvalidParentheses(s string) []string {
	var left, right = 0, 0
	var sol = map[string]bool{}

	if !strings.Contains(s, "(") && !strings.Contains(s, ")") {
		return []string{s}
	} else if strings.Contains(s, ")") && !strings.Contains(s, "(") {
		s = strings.Trim(s, ")")
		return []string{s}
	} else if strings.Contains(s, "(") && !strings.Contains(s, ")") {
		s = strings.Trim(s, "(")
		return []string{s}
	} else if len(s) == 3 {
		countClosed := 0
		for c:=0; c < len(s); c++ {
			if s[c] == ')' {
				countClosed++
			}
			if c == 1 && countClosed == 2 {
				return []string{""}
			}
		}
		if strings.Contains(s, ")(") {
			s = strings.Replace(s, ")(", "", 1)
			return []string{s}
		}
		return []string{"()"}
	}

	for _, c := range s {
		if c == '(' {
			left++
		} else if c == ')' {
			if left == 0 {
				right++
			} else {
				left--
			}
		}
	}

	var dfs = func(depth, left, right, left_rem, right_rem int, curr string){}

	dfs = func(depth, left, right, left_rem, right_rem int, curr string) {
		if depth == len(s) {
			if left == right && left_rem ==  0 && right_rem == 0 {
				sol[curr] = true
				return
			}
		} else {
			if s[depth] == '(' {
				if left_rem > 0 {
					dfs(depth + 1, left, right, left_rem - 1, right_rem, curr)
				}
				dfs(depth + 1, left + 1, right, left_rem, right_rem, curr + "(")
			} else if s[depth] == ')' {
				if right_rem > 0 {
					dfs(depth + 1, left, right, left_rem, right_rem - 1, curr)
				}
				if left > right {
					dfs(depth + 1, left, right + 1, left_rem, right_rem, curr + ")")
				}
			} else {
				dfs(depth + 1, left ,right, left_rem, right_rem, curr + string(s[depth]))
			}
		}
	}

	dfs(0,0,0, left, right, "")

	var res []string
	for k := range sol {
		if len(k) > 1 {
			res = append(res, k)
		}
	}
	if len(res) == 0 {
		return []string{""}
	}
	return res
}



