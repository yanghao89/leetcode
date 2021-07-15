package hashmap

import (
	"fmt"
	"sort"
	"testing"
)

func halfQuestions(questions []int) (ans int) {
	//maps := make(map[int]int, 0)
	//for _, v := range questions {
	//	maps[v]++
	//}
	//a := make([]int, 0, len(maps))
	//for _, v := range maps {
	//	a = append(a, v)
	//}
	sort.Sort(sort.Reverse(sort.IntSlice(questions)))
	fmt.Println(questions)
	//n := len(questions) / 2
	//for n > 0 {
	//	n -= a[ans]
	//	ans++
	//}
	//return
	return
}

func halfQuestions02(questions []int) (ans int) {
	ints := make([]int, 1001, 1001)
	for _, v := range questions {
		ints[v]++
	}
	sort.Ints(ints)
	for i, n := 1000, len(questions)/2; n > 0; i-- {
		ans++
		n -= ints[i]
	}
	return
}
func TestHalfQuestions(t *testing.T) {
	fmt.Println(halfQuestions([]int{1, 5, 1, 3, 4, 5, 2, 5, 3, 3, 8, 6}))
	//fmt.Println(halfQuestions02([]int{1, 5, 1, 3, 4, 5, 2, 5, 3, 3, 8, 6}))
}
