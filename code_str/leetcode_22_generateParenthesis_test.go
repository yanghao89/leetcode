package code_str

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) (ans []string) {
	if n == 0 {
		return
	}
	var (
		dfs = func(s string, L, R int) {}
	)
	dfs = func(s string, L, R int) {
		if L == 0 && R == 0 {
			ans = append(ans, s)
			return
		}
		if L > R {
			return
		}
		if L > 0 {
			dfs(fmt.Sprintf("%s(", s), L-1, R)
		}
		if R > 0 {
			dfs(fmt.Sprintf("%s)", s), L, R-1)
		}
	}
	dfs("", n, n)
	return
}

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
}
