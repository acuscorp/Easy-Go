package my_reflection

import (
	"fmt"
	"reflect"
)

func Init_reflection_examples() {

	checkingNilValues()

}

type StringSlice []string

func reflectionValues() {
	s := StringSlice{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	s2 := sv.Interface().(StringSlice)

	s2.forEach(myPrint)
}

func myPrint(d string) {
	print(d)
}

func (sl StringSlice) forEach(f func(string)) {
	for i := 0; i < len(sl); i++ {
		f(sl[i])
	}
}

// Use a reflection to check if an interface is nil
func checkingNilValues() {
	var a interface{}
	fmt.Printf("The 'a' variable has no value %v?\n\r ", hasNoValue(a))

	var c interface{} = a

	fmt.Printf("The 'c' structure variable has no value %v?\n ", hasNoValue(c))

	b := struct {
		aa int
		bb int
	}{
		aa: 1,
		bb: 2,
	}

	fmt.Printf("The 'b' structure variable has no value %v?\n ", hasNoValue(b))

}
func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}
func reflectionWithTags() {
	type Noe string
	type Foo struct {
		A int    `myTag:"value1"`
		B string `myTag:"value2"`
		C Noe    `myTag:"value2"`
	}

	var f Foo

	ft := reflect.TypeOf(f)

	for i := 0; i < ft.NumField(); i++ {
		curFiel := ft.Field(i)
		println(curFiel.Name, curFiel.Type.Name(), curFiel.Tag.Get("myTag"))
		println(curFiel.Name, curFiel.Type.Kind().String(), curFiel.Tag.Get("myTag"))
	}
}
func simpleReflection() {
	type Foo struct{}
	type Age int

	println("working with type Age")
	var x Age
	xtype := reflect.TypeOf(x)
	fmt.Println(xtype)
	fmt.Println(xtype.Name())
	fmt.Println(xtype.Kind())

	f := Foo{}
	println("\nworking with the structure\n")
	fType := reflect.TypeOf(f)
	fmt.Println(fType)
	fmt.Println(fType.Name())

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Elem().Name())
	fmt.Println(xpt.Elem().Kind())
}
