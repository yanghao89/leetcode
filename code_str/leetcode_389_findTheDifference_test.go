package code_str

import (
	"fmt"
	"sync"
	"testing"
)

func findTheDifference01(s, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}

func findTheDifference02(s, t string) byte {
	cnt := [26]int{}
	for _, v := range s {
		cnt[v-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return ch
		}
	}
}
func findTheDifference03(s, t string) byte {
	sum := 0
	for _, v := range s {
		sum -= int(v)
	}
	for _, v := range t {
		sum += int(v)
	}
	return byte(sum)
}

func TestFindTheDifference(t *testing.T) {
	var (
		wg sync.WaitGroup
	)
	ch := make(chan byte)
	strs := make([][]string, 3)
	str1 := []string{"abcd", "abcde"}
	strs[0] = append(strs[0], str1...)
	strs[1] = append(strs[1], str1...)
	strs[2] = append(strs[2], str1...)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < len(strs); i++ {
			ch <- findTheDifference02(strs[i][0], strs[i][1])
		}
	}()
	go func() {
		wg.Done()
		for {
			vv, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(string(vv))
		}
	}()
	wg.Wait()
}
