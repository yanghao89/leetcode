package hashmap

import "sort"

type TimeMap struct {
	maps map[string][]pair
}

func ConstructorTimeMap() TimeMap {
	return TimeMap{
		maps: make(map[string][]pair, 0),
	}
}

type pair struct {
	Value     string
	Timestamp int
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.maps[key] = append(this.maps[key], pair{Value: value, Timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs := this.maps[key]
	i := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].Timestamp > timestamp
	})
	if i > 0 {
		return pairs[i-1].Value
	}
	return ""
}
