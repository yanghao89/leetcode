package leetcode_28

func IsValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	maps := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for k := range s {
		//检查maps中存在对应的byte
		if maps[s[k]] > 0 {
			//如果栈的数量等于0 | 当前的byte 和栈内的byte 不相等直接false
			if len(stack) == 0 || stack[len(stack)-1] != maps[s[k]] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]
		} else {
			//压栈
			stack = append(stack, s[k])
		}
	}
	return len(stack) == 0
}
