package golikekotlin

import (
	"fmt"
	"reflect"
  "strings"
)
func MyTests {
    s := StringSlice{"a","b","c"}
  sv := reflect.ValueOf(s)
  s2 := sv.Interface().(StringSlice)

  s2.forEach(func(d string){
    print(d)
  })
  
  println()
  ms2 := s2.myMap(func(d string) string{
    return strings.ToUpper(d)
  })
  
  ms2.forEach(myPrint)
}

func myPrint(d string) {
  print(d)
}

func (sl StringSlice) forEach(f func(string)) {
  for i := 0; i < len(sl) ; i++ {
    f(sl[i])
  }
}

func (sl StringSlice) myMap (f func(string) string) StringSlice {
  newStringSlice := make(StringSlice,len(sl))
  for i := 0; i < len(sl) ; i++ {
     newStringSlice[i] = f(sl[i])
  }
  return newStringSlice
