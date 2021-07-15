package hashmap

func expectNumber(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	maps := make(map[int]struct{}, 0)
	for _, v := range scores {
		maps[v] = struct{}{}
	}
	return len(maps)
}
