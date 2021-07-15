package medium

import (
	"fmt"
	"sort"
	"testing"
)

var pair = []struct {
	value int
	text  string
}{
	{1, "I"},
	{4, "IV"},
	{5, "V"},
	{9, "IX"},
	{10, "X"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
}

func intToRoman(num int) (ans string) {
	sort.Slice(pair, func(i, j int) bool {
		return pair[i].value > pair[j].value
	})
	roman := make([]byte, 0)
	for _, v := range pair {
		for num >= v.value {
			num -= v.value
			roman = append(roman, v.text...)
		}
		if num == 0 {
			break
		}
	}
	ans = string(roman)
	return
}

func TestIntToRoman(t *testing.T) {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
}
