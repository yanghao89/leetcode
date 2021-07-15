package hashmap

import (
	"fmt"
	"testing"
)

type OrderedStream struct {
	data map[int]string
	max  int
	ptr  int
}

func ConstructorStream(n int) OrderedStream {
	return OrderedStream{
		data: make(map[int]string),
		max:  n,
		ptr:  1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) (ans []string) {
	this.data[idKey] = value
	if idKey == this.ptr {
		for i := idKey; i <= this.max; i++ {
			if val, ok := this.data[i]; ok {
				ans = append(ans, val)
			} else {
				this.ptr = i
				break
			}
		}
	}
	return
}

func TestOrderedStream(t *testing.T) {
	os := ConstructorStream(5)
	fmt.Println(os.Insert(3, "ccccc"))
	fmt.Println(os.Insert(1, "aaaaa"))
	fmt.Println(os.Insert(2, "bbbbb"))
	fmt.Println(os.Insert(5, "eeeee"))
	fmt.Println(os.Insert(4, "ddddd"))

}
