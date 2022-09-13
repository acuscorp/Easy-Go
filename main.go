package main
import (  
  shape "acuscorp.com/shape"
)

func main() {
    rec := shape.Rectangle{
      h: 2.0,
      w: 3.0,
    }
  cir := shape.Circle{
    r:10,
  }
  shape.printShapeArea(rec)
  shape.printShapeArea(cir)
  shape.whatShapeTypeIs(cir)
}


