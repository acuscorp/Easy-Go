package main
import (  
  shape "acuscorp.com/shape"
)

func main() {
    rec := shape.Rectangle{
      H: 2.0,
      W: 3.0,
    }
  cir := shape.Circle{
    R:10,
  }
  shape.PrintShapeArea(rec)
  shape.PrintShapeArea(cir)
  shape.WhatShapeTypeIs(cir)
}


