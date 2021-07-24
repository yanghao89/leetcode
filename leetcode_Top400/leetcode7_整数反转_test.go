package leetcode_Top400

import (
	"fmt"
	"math"
	"testing"
)

func reverse(x int) (ans int) {
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 {
			return 0
		}
		// 弹出 x 的末尾数字 digit
		digit := x % 10
		x /= 10
		// 将数字 digit 推入 rev 末尾
		ans = ans*10 + digit
	}
	return
}

func TestReverse(t *testing.T) {
	fmt.Println(reverse(123))
}
