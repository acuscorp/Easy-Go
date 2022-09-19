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

	res := processAndGather(
		func(v int) int {
			return v * v
		}, 10)

	fmt.Println(res)
}

func processAndGather(processor func(int) int, num int) []int {
	out := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(v int) {
			defer wg.Done()

			out <- processor(v)

		}(i)
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
