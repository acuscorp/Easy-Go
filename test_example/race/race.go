package race

import (
  "sync"
)

func GetCounter() int {
  var counter int
  var wg sync.WaitGroup
  wg.Add(5)
  for i:= 0; i < 5; i++ {
    go func(){
      for j:= 0; j < 10000; j++ {
        counter++
      }
      wg.Done()
    }()
  }
  wg.Wait()
  return counter
}