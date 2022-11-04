package main

import "fmt"

func main() {
	director := new(Director)
	Client(director)
}

func Client(director *Director) {
	builder := ConcreteBuilder1{}
	director.setBuilder(builder)
	product := director.makeProduct(1)
	fmt.Println(product.getName())

	product = director.makeProduct(0)
	fmt.Println(product.getName())
}

type Director struct {
	builder IBuilder
}

func (d *Director) setBuilder(builder IBuilder) {
	d.builder = builder
}

func (d Director) makeProduct(option int) IProduct {

	if option == 0 {
		return d.builder.buildProduct1()
	}
	return d.builder.buildProduct2()

}

type IBuilder interface {
	buildProduct1() IProduct
	buildProduct2() IProduct
}

type ConcreteBuilder1 struct{}

func (c ConcreteBuilder1) buildProduct1() IProduct {
	return concreteProduct1{}
}
func (c ConcreteBuilder1) buildProduct2() IProduct {
	return concreteProduct2{}
}

type ConcreteBuilder2 struct{}

func (c ConcreteBuilder2) buildProduct1() IProduct {
	return concreteProduct1{}
}
func (c ConcreteBuilder2) buildProduct2() IProduct {
	return concreteProduct2{}
}

type IProduct interface {
	getName() string
}

type concreteProduct1 struct{}

func (c concreteProduct1) getName() string {
	return "concrete Product 1"
}

type concreteProduct2 struct{}

func (c concreteProduct2) getName() string {
	return "concrete Product 2"
}
