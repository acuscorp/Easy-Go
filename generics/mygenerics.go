package generics

import "fmt"

func InitMyGenerics() {
	var myStack = Stack[int]{
		vals: []int{10, 20, 30},
	}

	myStack.Push(40)
	myStack.Push(50)
	// take last one and print it

	x, ok := myStack.Pop()
	fmt.Println(x, ok)

	fmt.Printf("myStack.Contains(50)?: %v\n", myStack.Contains(50))

	fmt.Println("adding 50 to myStack")
	myStack.Push(50)

	fmt.Printf("myStack.Contains(50)?: %v\n", myStack.Contains(50))

}

// if you don't use comparable, code will not be able to use == (equals to)
// when you use 'any' is just for putting and getting generic data
type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(v T) bool {
	for _, stackVal := range s.vals {
		if stackVal == v {
			return true
		}
	}
	return false
}
