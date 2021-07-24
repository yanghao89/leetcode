package leetcode_Top400

import (
	"fmt"
	"strings"
	"testing"
)

/**
1. 以 V 字型为一个循环, 循环周期为 n = (2 * numRows - 2) （2倍行数 - 头尾2个）
2. 对于字符串索引值 i 计算 x = i % n 确定在循环周期中的位置。
3. y = min(x, n -x)
*/
func convert(s string, numRows int) (ans string) {
	row := make([]string, numRows)
	n := 2*numRows - 2
	for i, v := range s {
		x := i % n
		row[min(x, n-x)] += string(v)
	}
	return strings.Join(row, "")
}

func TestConvert(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
