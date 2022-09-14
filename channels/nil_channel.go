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
  t := time.NewTimer(3 * time.Second)
  for {
    select{
      case n := <- ch:
        fmt.Println(n)
      case <- t.C:
        ch = nil
        fmt.Println("Stopped on reader")
    }
  }
}

func Writer(ch chan int){
  stopper := time.NewTimer(1 * time.Second)
  restarter := time.NewTimer(2 * time.Second)
  backupCh := ch
  for {
    select{
      case ch <- rand.Intn(40):
      case <- stopper.C:
        ch = nil       
        fmt.Println("Stopped on writer")
      case  <- restarter.C:
        ch = backupCh
        fmt.Println("restarted on writer")
    }
  }
}