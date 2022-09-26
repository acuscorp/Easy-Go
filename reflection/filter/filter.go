package filter

import (
	"fmt"
	"reflect"
)

func InFilter() {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	longNames := Filter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)

	fmt.Println(longNames)

	ages := []int{20, 50, 13}

	adults := Filter(ages, func(age int) bool {
		return age >= 18
	}).([]int)

	fmt.Println(adults)

	adults2 := myFilter(ages, func(a int) bool {
		return a >= 18
	})

	fmt.Printf("adults using generic filter: %v\n", adults2)
}

func Filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)

	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}

func myFilter[T any](slice []T, f func(T) bool) []T {
	var n []T

	for _, v := range slice {
		if f(v) {
			n = append(n, v)
		}
	}
	return n
}
