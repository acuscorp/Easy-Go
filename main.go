package main
import "fmt"

func main() {
    rec := Rectangle{
      h: 2.0,
      w: 3.0,
    }
  cir := Circle{
    r:10,
  }
  printShapeArea(rec)
  printShapeArea(cir)
}


type Shape interface{
  area() float32
}

type Rectangle struct {
  h float32
  w float32
}

type Circle struct {
  r float32
}

func (rec Rectangle) area() float32 {
  return rec.h * rec.w
}

func (cir Circle) area() float32 {
  return 3.1416 * cir.r
}

func printShapeArea(shape Shape) {
  area := shape.area()
  fmt.Println("the area is ", area)
}