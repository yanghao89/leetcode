package hashmap

import (
	"fmt"
	"reflect"
	"testing"
)

func findWords(words []string) (ans []string) {
	if len(words) == 0 {
		return
	}
	maps := map[byte]byte{
		'q': 0, 'w': 0, 'e': 0, 'r': 0, 't': 0, 'y': 0, 'u': 0, 'i': 0, 'o': 0, 'p': 0,
		'Q': 0, 'W': 0, 'E': 0, 'R': 0, 'T': 0, 'Y': 0, 'U': 0, 'I': 0, 'O': 0, 'P': 0,
		'a': 1, 's': 1, 'd': 1, 'f': 1, 'g': 1, 'h': 1, 'j': 1, 'k': 1, 'l': 1,
		'A': 1, 'S': 1, 'D': 1, 'F': 1, 'G': 1, 'H': 1, 'J': 1, 'K': 1, 'L': 1,
		'z': 2, 'x': 2, 'c': 2, 'v': 2, 'b': 2, 'n': 2, 'm': 2,
		'Z': 2, 'X': 2, 'C': 2, 'V': 2, 'B': 2, 'N': 2, 'M': 2,
	}
	for _, work := range words {
		isSameLine := true
		for i := 0; i < len(work)-1; i++ {
			if maps[work[i]] != maps[work[i+1]] {
				isSameLine = false
				break
			}
		}
		if isSameLine {
			ans = append(ans, work)
		}
	}
	return
}

type Value struct {
	Age int
}

func TestFindWords(t *testing.T) {
	v1 := Value{Age: 10}
	v2 := Value{Age: 10}
	fmt.Println(reflect.DeepEqual(v1, v2))
	fmt.Println(&v1 == &v2)

	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
