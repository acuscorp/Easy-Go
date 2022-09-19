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
