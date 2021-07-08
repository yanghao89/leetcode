package hashmap

import (
	"fmt"
	"math"
	"testing"
)

func findRestaurant(list1, list2 []string) (ans []string) {
	maps := make(map[int][]string, 0)
	for i := range list1 {
		for j := range list2 {
			if list1[i] == list2[j] {
				maps[i+j] = append(maps[i+j], list1[i])
			}
		}
	}
	minIndexSum := math.MaxInt64
	for k := range maps {
		minIndexSum = int(math.Min(float64(minIndexSum), float64(k)))
	}
	ans = maps[minIndexSum]
	return
}

func TestFindRestaurant(t *testing.T) {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}))
}
