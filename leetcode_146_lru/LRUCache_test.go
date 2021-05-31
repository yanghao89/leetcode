package leetcode_146_lru

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	lru := Constructor(1)

	lru.Put(10, 1)
	lru.Put(9, 2)
	lru.Put(8, 3)

	fmt.Println(lru.Get(3))
}
