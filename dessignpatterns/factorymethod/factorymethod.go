package main

import "fmt"

func main() {
	Client(Factory1{})
}
func Client(f IFactory) {

	product1 := f.CrateProduct1(P1) // casting to Product1 otherwise it will return IProduct

	fmt.Println(product1.GetName())
}

type IProduct interface {
	GetName() string
}
type IFactory interface {
	CrateProduct1(Product) IProduct
}

type Factory1 struct {
	Name string
}

type Product int

const (
	Undefined Product = iota
	P1
	P2
	P3
)

func (f Factory1) CrateProduct1(p Product) IProduct {
	var product IProduct
	switch p {
	case P1:
		product = Product1{
			Name: "prduct 1",
		}
	case P2:
		product = Product2{
			Name: "prduct 2",
		}
	case P3:
		product = Product3{}
	}

	return product
}

type Product1 struct {
	Name string
	Age  int
}

type Product2 struct {
	Name string
}

type Product3 struct {
}

func (p Product1) GetName() string {
	return p.Name
}
func (p Product2) GetName() string {
	return p.Name
}

func (p Product3) GetName() string {
	return "no name for product3"
}
