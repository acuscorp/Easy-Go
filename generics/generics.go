package generics

import (
	"fmt"
	"math"
)

func InitGenerics() {
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer := FindCloser(pair2Da, pair2Db)
	fmt.Println(closer)
}

type Point2D struct {
	X, Y int
}

func (p2 Point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.X, p2.Y)
}

func (p2 Point2D) Diff(from Point2D) float64 {
	x := p2.X - from.X
	y := p2.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Pair[T any] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)

	if d1 < d2 {
		return pair1
	}
	return pair2
}
