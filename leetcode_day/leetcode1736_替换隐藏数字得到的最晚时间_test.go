package day

import (
	"fmt"
	"testing"
)

func maximumTime(time string) string {
	bt := []byte(time)
	if bt[0] == '?' {
		if '4' <= bt[1] && bt[1] <= '9' {
			bt[0] = '1'
		} else {
			bt[0] = '2'
		}
	}
	if bt[1] == '?' {
		if bt[0] == '2' {
			bt[1] = '3'
		} else {
			bt[1] = '9'
		}
	}
	if bt[3] == '?' {
		bt[3] = '5'
	}
	if bt[4] == '?' {
		bt[4] = '9'
	}
	return string(bt)
}

func TestMaximumTime(t *testing.T) {
	fmt.Println(maximumTime("2?:?0"))
}
