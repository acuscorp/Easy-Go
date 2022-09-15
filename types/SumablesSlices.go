package types

import "fmt"

type SumableSlice []int

func (s SumableSlice) sum() int {
  sum := 0
  for _ ,v := range s {
    sum += v
  }
  return sum
}
func InitSumables() {
  var s SumableSlice = SumableSlice{1,1,2,3,4,5}
  
  fmt.Println("the sum is ", s.sum())  
}