package hashmap

import (
	"fmt"
	"testing"
)

func isHappy(n int) bool {
	var (
		stup func(n int) int
	)
	stup = func(n int) int {
		num := 0
		for n > 0 {
			num += (n % 10) * (n % 10)
			n = n / 10
		}
		return num
	}
	m := make(map[int]bool, 0)
	for ; n != 1 && !m[n]; n, m[n] = stup(n), true {

	}
	return n == 1

}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}
