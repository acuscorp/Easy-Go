package channels

import (
  "fmt"
  "time"
  "math/rand"
  "sync/atomic"
)

var running int64 =  0

func work(){
  atomic.AddInt64(&running, 1)
  fmt.Printf("[%d", running)
  time.Sleep( time.Duration(rand.Intn(10)) * time.Second)
  atomic.AddInt64(&running, -1)
  fmt.Printf("]")
}
func worker(semaphore chan bool){
  <- semaphore
  work()
  semaphore <- true
  }

func IniBufferedCahnnel(){
  semaphore := make(chan bool, 10)

  for i := 0; i < 1000; i++ {
    go worker(semaphore)
  }

  for i:= 0; i<cap(semaphore); i++ {
    semaphore <- true
  }

  time.Sleep(30 *time.Second)

}