package shape

import "fmt"
func initialize(){
      rec := Rectangle{
      H: 2.0,
      W: 3.0,
    }
  cir := Circle{
    R:10,
  }
  PrintShapeArea(rec)
  PrintShapeArea(cir)
  WhatShapeTypeIs(cir)
}
func WhatShapeTypeIs(i interface{}) {
  switch v := i.(type) {
  	case Rectangle:
  		fmt.Printf("This is a %T and has an area %v", v,v.Area())
  	case Circle:
  		fmt.Printf("This is a %T and has an area %v", v,v.Area())
  	default:
  		fmt.Printf("I don't know about type %T!\n", v)
	}
}



type Shape interface{
  Area() float32
}

type Rectangle struct {
  H float32
  W float32
}

type Circle struct {
  R float32
}

func (rec Rectangle) Area() float32 {
  return rec.H * rec.W
}

func (cir Circle) Area() float32 {
  return 3.1416 * cir.R
}

func PrintShapeArea(shape Shape) {
  area := shape.Area()
  fmt.Println("the area is ", area)
}