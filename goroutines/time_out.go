package goroutines 

import (
  "fmt"
  "time"
  "math/rand"
  )

func InitTimeOut(){
  fmt.Println("starting")
  done := make(chan struct{})
  
  go func(){
    fmt.Println("Getting data from a servet...")
    rand.Seed(time.Now().UnixNano())
    var timeToSleep = time.Duration(rand.Intn(10))
    fmt.Println("this will take" ,timeToSleep*time.Second, " to response")  
    time.Sleep(timeToSleep * time.Second)
    close(done)
  }()
    select {
      case <- done:
        fmt.Println("Sucessful request")
      case <- time.After(3 * time.Second):
        fmt.Println("It is much time, need to finish")
    }
  
}