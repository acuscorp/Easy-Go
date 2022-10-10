package typos

import "fmt"

type AType struct {
	name string
}

func (a AType) GetName() string {
	return a.name
}
func (a *AType) SetName(name string) {
	a.name = name
}

type BType struct {
	V string
}

var b BType

func (b BType) SetB(a BType) {
	b = a
}
func (a AType) GetB() BType {
	return b
}
func InitTypos(n BType) {
	fmt.Println(n.V)

}
