package code_str

import (
	"fmt"
	"testing"
)

func StrToArray(str string) []string {
	slow, fast := 0, 0
	res := []string{}
	for fast <= len(str)-1 {
		if fast == len(res) || str[slow] != str[fast] {
			if fast-slow > 1 {
				res = append(res, str[slow:fast])
			}
			slow = fast
		}
		fast++
	}
	return res
}

func TestStrToArray(t *testing.T) {
	fmt.Println(StrToArray("aabbcdde"))
	fmt.Println(StrToArray("abcabcabcabc"))
	fmt.Println(StrToArray("bcbdddttyywwiiilllammnn"))
}
