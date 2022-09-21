package contexts

import (
  "context"
  "time"
  "fmt"
)
func InitTimers(){
  contextWithTimeOut()
}
// Cancelling child context when parent context is cancelled
func contextWithTimeOut(){
  ctx := context.Background()
  parent, cancel := context.WithTimeout(ctx, 2 * time.Second)  // after 2 seconds will cancel. This is the parent context
  defer cancel()  // cancel when finish the local function

  child, cancel2 := context.WithTimeout(parent, 3 * time.Second) // cancel after 3 seconds. This is a child context
  defer cancel2() // cancel when finish the local function
  start := time.Now()
  <- child.Done() // wait to receive Done() on child channel
  end := time.Now()
  fmt.Println(end.Sub(start))
  
  
}