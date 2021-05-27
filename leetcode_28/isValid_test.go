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
