package leetcode_28

import (
	"fmt"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
	ch = make(chan string)
)

func valID() {
	defer wg.Done()
	array := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "{}({}[]{{）", "{}[]{([}{}}", "(("}
	for i := 0; i < len(array); i++ {
		check := IsValid(array[i])
		if check {
			ch <- array[i]
		}
	}
	close(ch)
}

func IsValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	maps := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for k := range s {
		//检查maps中存在对应的byte
		if maps[s[k]] > 0 {
			//如果栈的数量等于0 | 当前的byte 和栈内的byte 不相等直接false
			if len(stack) == 0 || stack[len(stack)-1] != maps[s[k]] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]
		} else {
			//压栈
			stack = append(stack, s[k])
		}
	}
	return len(stack) == 0
}

func TestIsValid(t *testing.T) {
	wg.Add(2)
	go valID()
	go func() {
		defer wg.Done()
		for {
			v, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(v)
		}
	}()
	wg.Wait()

}
