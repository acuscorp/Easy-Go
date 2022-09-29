package generics

import (
	"fmt"
)

func InitXGenerics() {
	var slice = Slice[int]{10, 11, 12, 13}
	slice2 := slice.XMap(func(v int) int {
		return v * v

	})
	for _, v := range slice {
		fmt.Printf("slice (no map) v -> %v,\n", v)
	}

	println()
	for _, v := range slice2 {
		fmt.Printf("slice mapped v*v->%v,\n", v)
	}

	reduced := slice.XReduce(0, func(a, b int) int {
		return a + b
	})

	fmt.Printf("sum of '10, 11, 12, 13' : %v\n", reduced)

	filtered := slice.XFilter(func(v int) bool {
		return v > 11
	})

	fmt.Printf("filtered values grater than 11 : %v\n", filtered)

	// using let
	var n1 Struct[int] = Struct[int]{Value: 35}
	nn := n1.Let(func(n int) int {
		return n + n
	})

	fmt.Printf("nn: %v\n", nn.Value)

	var n2 = Struct[string]{Value: "a"}
	nn2 := n2.Let(func(n string) string {
		return n + "b"
	})

	fmt.Printf("nn2: %v\n", nn2.Value)

	// using apply

	n1.Apply(func(s *Struct[int]) {
		s.Value += 10
	})

	fmt.Printf("applying 35 + 10 n: %v\n", n1.Value)

	n2.Apply(func(s *Struct[string]) {
		s.Value += ",b,c"
	})
	fmt.Printf("applying b,c to a n: %v\n", n2.Value)
}

type A[T any] interface {
	execute(t T)
}

func (a A[T]) Pu(b T) {
	a.execute(b)
}

func (s *Struct[T]) Apply(f func(*Struct[T])) {
	f(s)
}

type Slice[T any] []T

func (s Slice[T]) XMap(f func(T) T) []T {
	var slice = make(Slice[T], len(s))
	for i, v := range s {
		slice[i] = f(v)
	}

	return slice
}

func (s Slice[T]) XReduce(initializer T, f func(T, T) T) T {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func (s Slice[T]) XFilter(f func(v T) bool) Slice[T] {
	var sc = Slice[T]{}
	for _, v := range s {
		if f(v) {
			sc = append(sc, v)
		}
	}
	return sc
}

type Struct[T any] struct {
	Value T
}

func (n Struct[T]) Let(f func(T) T) Struct[T] {

	return Struct[T]{Value: f(n.Value)}
}
