package binary_search

import (
	"fmt"
	"math"
	"testing"
)

func arrangeCoins01(n int) int {
	L, R := 0, n
	for L < R {
		mid := L + (R - L) + 1>>1
		sum := (1 + mid) * mid / 2
		if sum > n {
			R = mid - 1
		} else {
			L = mid
		}
	}
	return L
}

// x = ( √(1+8n) - 1 ) / 2
func arrangeCoins02(n int) int {
	return int((math.Sqrt(float64(8*n+1)) - 1) / 2)
}

func TestArrangeCoins(t *testing.T) {
	fmt.Println(arrangeCoins01(5))
	fmt.Println(arrangeCoins02(1))
}
