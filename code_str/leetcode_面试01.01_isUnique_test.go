package code_str

import (
	"fmt"
	"testing"
)

func isUnique(astr string) (ans bool) {
	mark := 0
	for _, str := range astr {
		v := str - 'a'
		// 两个位都为1时，结果才为1
		if mark&(1<<v) != 0 {
			return
		} else {
			//两个位都为0时，结果才为0
			mark |= (1 << v)
		}
	}
	return true
}

func TestIsUnique(t *testing.T) {
	fmt.Println(isUnique("leetcode"))
	fmt.Println(isUnique("abc"))
}
