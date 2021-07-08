package hashmap

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func subdomainVisits(cpdomains []string) (ans []string) {
	count := make(map[string]int, 0)
	for _, v := range cpdomains {
		split := strings.Split(v, " ")
		num, _ := strconv.Atoi(split[0])
		for {
			//统计第一个访问域名
			count[split[1]] += num
			//找当前域名第一个点的位置
			dotIndex := strings.Index(split[1], ".")
			if dotIndex < 0 {
				break
			}
			//切割域名的
			split[1] = split[1][dotIndex+1:]
		}
	}
	ans = make([]string, 0, len(count))
	for k, v := range count {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return
}

func TestSubdomainVisits(t *testing.T) {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
}
