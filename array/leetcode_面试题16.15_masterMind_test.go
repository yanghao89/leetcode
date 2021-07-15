package array

import (
	"fmt"
	"testing"
)

func masterMind(solution string, guess string) (ans []int) {
	ans = make([]int, 2)
	maps := make(map[byte]int, 0)
	for i := 0; i < len(solution); i++ {
		if solution[i] == guess[i] {
			ans[0]++
		}
		maps[byte(solution[i])-'A']++
	}
	count := 0
	for i := 0; i < len(guess); i++ {
		if _, ok := maps[byte(guess[i])-'A']; ok && maps[byte(guess[i])-'A'] > 0 {
			count++
			maps[byte(guess[i])-'A']--
		}
	}
	ans[1] = count - ans[0]
	return ans
}

func masterMind02(solution string, guess string) (ans []int) {
	maps := make([]int, 26, 26)
	total, count := 0, 0
	for i := 0; i < 4; i++ {
		maps[int(guess[i]-'A')]++
	}
	for i := 0; i < 4; i++ {
		if guess[i] == solution[i] {
			count++
		}
		if maps[int(solution[i]-'A')] > 0 {
			total++
			maps[int(solution[i]-'A')]--
		}
	}
	ans = append(ans, count, total-count)
	return
}

func TestMasterMind(t *testing.T) {
	//fmt.Println(masterMind("RGBY", "GGRR"))
	fmt.Println(masterMind02("RGBY", "GGRR"))
}
