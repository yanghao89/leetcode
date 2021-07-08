package binary_search

import (
	"fmt"
	"math"
	"testing"
)

func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}

func TestMySqrt(t *testing.T) {
	fmt.Println(mySqrt(4))
}
