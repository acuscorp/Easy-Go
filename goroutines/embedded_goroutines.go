package goroutines

import (
  "fmt"
  "time"
  )
// printNumbers returns a just write channel and a function to close the goroutine

// the sanw task can be done using <- chan int (only read) but the function rather to print the numbers it would write the numbers and in other function could be printed or read
func printNumbers() (chan <- int, func()) {
	ch := make(chan int)
	done := make(chan interface{})

	close := func() {
		close(done)
	}
  
	go func() {
		for {
			select {
			case i := <-ch:
				fmt.Println("Number: ", i)
			case <-done:
				return
			}
		}
	}()

	return ch, close
}

func InitEmbeddedFunction() {
	ch, closePrintNumbers := printNumbers()

	for i := 0; ; i++ {
		
    time.Sleep(1 * time.Second)
		ch <- i
    if i == 20 {
			break
		}
	}
  closePrintNumbers()
}
