package hashmap

import (
	"fmt"
	"testing"
)

func numEquivDominoPairs01(dominoes [][]int) (ans int) {
	cnt := [100]int{}
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return
}
func numEquivDominoPairs02(dominoes [][]int) (ans int) {
	dm := make(map[string]int, 0)
	for _, d := range dominoes {
		key := hash(d)
		if v, ok := dm[key]; ok {
			dm[key] = v + 1
		} else {
			dm[key] = 1
		}
	}
	for _, v := range dm {
		ans += v * (v - 1) / 2
	}
	return
}
func hash(d []int) string {
	if d[0] > d[1] {
		return fmt.Sprintf("%d,%d", d[0], d[1])
	}
	return fmt.Sprintf("%d,%d", d[1], d[0])
}

func TestNumEquivDominoPairs(t *testing.T) {
	fmt.Println(numEquivDominoPairs01([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
	fmt.Println(numEquivDominoPairs02([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}
