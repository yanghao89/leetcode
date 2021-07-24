package main

import (
	"strconv"
	"testing"
)

func BenchmarkToString(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ToString(i)
	}
	b.StopTimer()
}

func ToString(i int) string {
	return strconv.Itoa(i)
}
