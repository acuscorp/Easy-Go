package main

import "fmt"

func main() {
	var cFactory1 = ConcreteFactory1{}
	Client(cFactory1)
}

func Client(factory AbstractFactory) {
	productA := factory.createProductA()
	fmt.Println(productA.name)
}

type ProductA struct {
	name string
}
type ProductB struct {
	name string
}

type AbstractFactory interface {
	createProductA() ProductA
	createProductB() ProductB
}

type ConcreteFactory1 struct{}

func (cFactory1 ConcreteFactory1) createProductA() ProductA {
	return ProductA{
		name: "ConcreteFactory1: PrductA",
	}
}
func (cFactory1 ConcreteFactory1) createProductB() ProductB {
	return ProductB{
		name: "ConcreteFactory1: PrductB",
	}
}

type ConcreteFactory2 struct{}

func (cFactory2 ConcreteFactory2) createProductA() ProductA {
	return ProductA{
		name: "ConcreteFactory2: ProductA",
	}
}
func (cFactory2 ConcreteFactory2) createProductB() ProductB {
	return ProductB{
		name: "ConcreteFactory2: ProductB",
	}
}
