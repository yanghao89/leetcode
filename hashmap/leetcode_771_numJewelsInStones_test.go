package hashmap

import (
	"fmt"
	"testing"
)

func numJewelsInStones(jewels string, stones string) (ans int) {
	maps := make(map[byte]bool, 0)
	for i := range jewels {
		maps[jewels[i]] = true
	}
	for i := range stones {
		if maps[stones[i]] {
			ans++
		}
	}
	return
}

func TestNumJewelsInStones(t *testing.T) {
	strs := []string{"aA", "aAAbbbb", "z", "ZZ"}
	for i := 0; i < len(strs); i += 2 {
		fmt.Println(numJewelsInStones(strs[i], strs[i+1]))
	}

}
