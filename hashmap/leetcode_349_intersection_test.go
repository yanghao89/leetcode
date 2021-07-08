package hashmap

import (
	"fmt"
	"testing"
)

func intersection(num1, num2 []int) (ints []int) {
	m1 := make(map[int]struct{}, 0)
	for _, v := range num1 {
		m1[v] = struct{}{}
	}
	m2 := make(map[int]struct{}, 0)
	for _, v := range num2 {
		m2[v] = struct{}{}
	}
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	for v := range m1 {
		if _, has := m2[v]; has {
			ints = append(ints, v)
		}
	}
	return
}

func TestIntersection(t *testing.T) {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
