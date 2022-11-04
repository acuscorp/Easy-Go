package main

import "fmt"

func main() {
	cPrototype1 := ConcretePrototype1{
		pName: "I'm concrete prototype one",
	}

	cPrototype1.printName()
	cPrototypeClone := cPrototype1.clone()
	cPrototypeClone.setName("Cloned prototype one")
	cPrototypeClone.printName()
}

type Prototype interface {
	printName()
	setName(name string)
	clone() Prototype
}

type ConcretePrototype1 struct {
	pName string
}

func (p *ConcretePrototype1) clone() Prototype {
	return &ConcretePrototype1{}
}

func (p *ConcretePrototype1) printName() {
	fmt.Println(p.pName + " I'm doing something")
}

func (p *ConcretePrototype1) setName(name string) {
	p.pName = name
}
