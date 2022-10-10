package main

import (
	g "main/typos"
)

func main() {
	a := g.AType{}
	a.SetName("noe")

	b := a.GetB()     //AType has acces to BType
	b.V = a.GetName() // now you can work with BType, But what happend if
	g.InitTypos(b)    // you can't get values from variable
}
