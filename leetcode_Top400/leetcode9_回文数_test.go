package leetcode_Top400

import (
	"fmt"
	"testing"
)

func isPalindrome(n int) (ans bool) {
	if n < 0 || (n%10 == 0 && n != 0) {
		return
	}
	cur := 0
	for n > cur {
		cur = cur*10 + n%10
		n /= 10
	}
	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return n == cur || n == cur/10
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome(9898989))
}
