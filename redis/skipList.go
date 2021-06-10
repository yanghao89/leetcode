package redis

import "math/rand"

type SkipList struct {
}

const (
	SKIPLIST_P = 5
)

func GetRandomLevel() int {
	L := 1
	for (rand.Intn(0xFFFF) & 0xFFFF) < (0xFFFF >> SKIPLIST_P) {
		L += 1
	}
	return L
}

//func ()  {
//
//}
