package hashmap

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) (ans int) {
	maps := make(map[int]*Employee, 0)
	for _, v := range employees {
		maps[v.Id] = v
	}
	var (
		dfs func(int)
	)
	dfs = func(subID int) {
		employee := maps[subID]
		if nil != employee {
			ans += employee.Importance
			for _, subId := range employee.Subordinates {
				dfs(subId)
			}
		}
	}
	dfs(id)
	return
}

func TestGetImportance(t *testing.T) {
	slices := make([]*Employee, 0)
	slices = append(slices, &Employee{
		Id:           1,
		Importance:   5,
		Subordinates: []int{2, 3},
	})
	slices = append(slices, &Employee{
		Id:           2,
		Importance:   3,
		Subordinates: []int{4},
	})
	slices = append(slices, &Employee{
		Id:           3,
		Importance:   4,
		Subordinates: []int{},
	})
	slices = append(slices, &Employee{
		Id:           4,
		Importance:   1,
		Subordinates: []int{},
	})
	fmt.Println(getImportance(slices, 2))
}
