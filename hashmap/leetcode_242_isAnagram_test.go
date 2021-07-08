package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func isAnagram01(s string, t string) bool {
	s1, t1 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] > s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] > t1[j]
	})
	return string(s1) == string(t1)
}

func isAnagram02(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[rune]int, 0)
	for _, v := range s {
		maps[v]++
	}
	for _, v := range t {
		maps[v]--
		if maps[v] < 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram01("anagram", "nagaram"))
	fmt.Println(isAnagram01("rat", "car"))
	fmt.Println(isAnagram02("anagram", "nagaram"))
	fmt.Println(isAnagram02("rat", "car"))
	nums := []int{3, 0, 1}
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= i ^ nums[i]
	}
	fmt.Println(res)
}
