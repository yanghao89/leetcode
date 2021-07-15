package hashmap

import (
	"fmt"
	"testing"
)

func isPathCrossing(path string) bool {
	maps := make(map[string]int, 0)
	x, y := 0, 0
	maps[fmt.Sprintf("%d%d", x, y)] = 1
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			y += 1
		case 'S':
			y -= 1
		case 'W':
			x -= 1
		case 'E':
			x += 1
		}
		key := fmt.Sprintf("%d%d", x, y)
		if _, ok := maps[key]; !ok {
			maps[key] = 1
		} else {
			return true
		}
	}
	return false
}

func TestIsPathCrossing(t *testing.T) {
	fmt.Println(isPathCrossing("NNSWWEWSSESSWENNW"))
	fmt.Println(isPathCrossing("NEEEEEEEEEENNNNNNNNNNWWWWWWWWWW"))
}
