package utils

import (
	"reflect"
	"strings"
)

// StructToSlice 结构体转切片
func StructToSlice(inter interface{}) []interface{} {
	t := reflect.TypeOf(inter)
	values := make([]interface{}, 0)
	v := reflect.ValueOf(inter)
	for i := 0; i < t.NumField(); i++ {
		keys := t.Field(i)
		str := keys.Name[0]
		key0 := strings.ToLower(string(str))
		key := key0 + keys.Name[1:]
		value := v.Field(i).Interface()
		values = append(values, key, value)
	}
	return values
}
