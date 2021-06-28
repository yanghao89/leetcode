package hashmap

import (
	"fmt"
	"testing"
)

//
func isIsomorphic(s string, t string) bool {
	s2t, t2s := make(map[string]byte, 0), make(map[string]byte, 0)
	for i := range s {
		x, y := s[i], t[i]
		if (s2t[string(x)] > 0 && s2t[string(x)] != y) || (t2s[string(y)] > 0 && t2s[string(y)] != x) {
			return false
		}
		s2t[string(x)] = y
		t2s[string(y)] = x
	}
	return true
}

func TestIsIsomorphic(t *testing.T) {
	// edd -> agg  true
	fmt.Println(isIsomorphic("egg", "add"))
	// far -> boo  false
	fmt.Println(isIsomorphic("foo", "bar"))
	// tater -> piple  true
	fmt.Println(isIsomorphic("paper", "title"))
}
