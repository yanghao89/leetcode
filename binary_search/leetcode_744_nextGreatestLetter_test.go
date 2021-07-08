package binary_search

import (
	"fmt"
	"sync"
	"testing"
)

func nextGreatestLetter(letters []byte, target byte) byte {
	L, R := 0, len(letters)-1
	for L < R {
		mid := L + (R-L)>>1
		if letters[mid] <= target {
			L = mid + 1
		}
		if letters[mid] > target {
			R = mid
		}
	}
	if letters[L] > target {
		return letters[L]
	}
	return letters[0]
}
func TestNextGreatestLetter(t *testing.T) {
	var (
		wg sync.WaitGroup
	)
	wg.Add(5)
	ch := make(chan byte)
	go func() {
		wg.Done()
		ch <- nextGreatestLetter([]byte(`cfj`), byte('a'))
	}()
	go func() {
		wg.Done()
		ch <- nextGreatestLetter([]byte(`cfj`), byte('c'))

	}()
	go func() {
		wg.Done()
		ch <- nextGreatestLetter([]byte(`cfj`), byte('d'))
	}()
	go func() {
		wg.Done()
		ch <- nextGreatestLetter([]byte(`cfj`), byte('g'))
	}()
	go func() {
		wg.Done()
		for {
			select {
			case vv, ok := <-ch:
				if !ok {
					break
				}
				fmt.Println(string(vv))
				//close(ch)
			}
		}
	}()
	wg.Wait()
}
