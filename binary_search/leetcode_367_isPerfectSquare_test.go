package binary_search

import (
	"fmt"
	"testing"
)

func isPerfectSquare01(num int) bool {
	if num < 2 {
		return true
	}
	var (
		x, guessSquared int
	)
	L, R := 2, num/2
	for L <= R {
		x = L + (R-L)>>1
		guessSquared = x * x
		if guessSquared == num {
			return true
		}
		if guessSquared > num {
			R = x - 1
		} else {
			L = x + 1
		}
	}
	return false
}

func isPerfectSquare02(num int) bool {
	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}
	return r*r == num
}

func TestIsPerfectSquare(t *testing.T) {
	fmt.Println(isPerfectSquare01(14))
	fmt.Println(isPerfectSquare02(16))
}
