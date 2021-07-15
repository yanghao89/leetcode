package hashmap

import (
	"fmt"
	"testing"
)

func countBalls(lowLimit int, highLimit int) (ans int) {
	maps := make(map[int]int, 0)
	for i := lowLimit; i <= highLimit; i++ {
		tmp, box := i, 0
		for tmp != 0 {
			box += tmp % 10
			tmp /= 10
		}
		maps[box]++
		ans = max(maps[box], ans)
	}
	return
}

func TestCountBalls(t *testing.T) {
	fmt.Println(countBalls(1, 10))
	fmt.Println(countBalls(5, 15))
	fmt.Println(countBalls(19, 28))
}
