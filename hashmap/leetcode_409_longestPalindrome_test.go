package hashmap

import (
	"testing"
)

func longestPalindrome(s string) int {
	ints := make([]int, 128)
	for _, v := range s {
		ints[v]++
	}
	count := 0
	for _, v := range ints {
		count += v / 2 * 2
		if v%2 == 1 && count%2 == 0 {
			count++
		}
	}
	return count
}

func TestLongestPalindrome(t *testing.T) {
	//count := longestPalindrome("cbccacdcd")
	//fmt.Println(5 % 2)
	//var (
	//	wg sync.WaitGroup
	//)
	//wg.Add(2)
	//c := make(chan int)
	//go func() {
	//	wg.Done()
	//	c <- longestPalindrome("abccccdd")
	//	close(c)
	//}()
	//go func() {
	//	wg.Done()
	//	for {
	//		select {
	//		case v, ok := <-c:
	//			if !ok {
	//				break
	//			}
	//			fmt.Println(v)
	//		}
	//	}
	//}()
	//wg.Wait()
}
