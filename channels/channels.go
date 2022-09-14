package channels

import (
  "fmt"
  "time"
)

func printer(msg string, stopCh chan bool) {
  for {
    select{
      case <- stopCh:  // this is cordinated by this channel
        return
      default:
        fmt.Printf("%s\n", msg)
    }
  }
    
}

func Initialize(){
  goCh := make(chan bool)

  for i:= 0; i <10; i++ {
    go printer(fmt.Sprintf("ptinter: %d", i), goCh)
  }

  time.Sleep(5*time.Second)
  close(goCh)
  time.Sleep(5*time.Second)
}