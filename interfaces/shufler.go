package interfaces

import (
  "math/rand"
  "fmt"
)
type shuffler interface {
  Len() int
  Swap(i,j int)
}

func shuffle(s shuffler){
  for i := 0; i < s.Len(); i++ {
    j := rand.Intn(s.Len()-1)
    s.Swap(i, j)
  }
}

type intSlice []int 

func (i intSlice) Len() int {
  return len(i)
}

func (is intSlice) Swap(i,j int) {
  is[i], is[j] = is[j] ,is[i]
}

type stringSlice []string 

func (i stringSlice) Len() int {
  return len(i)
}

func (is stringSlice) Swap(i,j int) {
  is[i], is[j] = is[j] ,is[i]
}

func InitShuffle(){
  is := intSlice{1,2,3,4,5,6,7,8,9}
  shuffle(is)
  fmt.Printf("%q\n", is)

  ss := stringSlice{"a","b","c","d"}
  shuffle(ss)
  fmt.Printf("%s", ss)
  
}