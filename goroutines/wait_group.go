package goroutines

import (
	"fmt"
	"sync"
)

func InitWaitGroupr() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := range make([]int, 3) {
		go func(v int) {
			defer wg.Done()
			fmt.Println("func ", v)
		}(i)
	}
	wg.Wait()
}

func InitWaitGroupWithChannels() {
  ch := make(chan int,5)
  a := processAndGather(ch, func(v int) int{
    return v*v
  }, 5)
  fmt.Println(a)
  close(ch)
}

func processAndGather(in <- chan int, processor func(int) int, num int) []int {
  out := make(chan int, num)
  var wg sync.WaitGroup
  wg.Add(num)

  for i := 0; i < num; i++ {
    go func() {
      defer wg.Done()
      for v := range in {
        out <- processor(v)
      }
    }()
  }
  
  go func() {
    wg.Wait()
    close(out)
  }()

  var result []int

  for v := range out {
    result = append(result, v)
  }
  return result
}