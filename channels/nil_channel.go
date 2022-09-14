package channels

import (
  "fmt"
  "time"
  "math/rand"
)

func InitializeNilChannel(){
  ch := make(chan int)
  go Reader(ch)
  go Writer(ch)
  time.Sleep(10 * time.Second)
}

func Reader(ch chan int) {
  t := time.NewTicker(3 * time.Second)
  for {
    select{
      case n := <- ch:
        fmt.Println(n)
      case <- t.C:
        ch = nil
    }
  }
}

func Writer(ch chan int){
  for {
    ch <- rand.Intn(40)
  }
}