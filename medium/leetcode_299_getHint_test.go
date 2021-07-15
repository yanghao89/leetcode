package medium

import (
	"fmt"
	"testing"
)

func getHint(secret string, guess string) string {
	maps := make(map[byte]int, 0)
	x, y := 0, 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			x++
		} else {
			maps[secret[i]]++
			maps[guess[i]]--
		}
	}
	y = len(secret) - x
	for _, v := range maps {
		if v > 0 {
			y -= v
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}

func TestGetHint(t *testing.T) {
	fmt.Println(getHint("1807", "7810"))
}
