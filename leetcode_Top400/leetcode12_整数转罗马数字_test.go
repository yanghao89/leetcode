package leetcode_Top400

import (
	"fmt"
	"testing"
)

var (
	p = []struct {
		Value int
		Str   string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
)

func intToRoman(num int) string {
	roman := make([]byte, 0)
	for _, s := range p {
		for num >= s.Value {
			num -= s.Value
			roman = append(roman, s.Str...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}

func TestIntToRoman(t *testing.T) {
	fmt.Println(intToRoman(3))
}
