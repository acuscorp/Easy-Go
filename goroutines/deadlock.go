package goroutines
import "fmt"
func InitDeadLock(){
  ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
  
  
}

func CloaureWrongOper(){
  arr := []int{1,2,3,4,5}
  ch := make(chan int)
  for _, v := range arr {
    go func(){
      ch <- v * 2
    }()
  }

  for i:=0; i < len(arr); i++{
    fmt.Println(<-ch)
  }
}

func CloaureRightOperOne(){
  arr := []int{1,2,3,4,5}
  ch := make(chan int)
  for _, v := range arr {
    v:= v// using shadowing
    go func(){
      ch <- v * 2
    }()
  }

  for i:=0; i < len(arr); i++{
    fmt.Println(<-ch)
  }
}

func CloaureRightOperTwo(){
  arr := []int{1,2,3,4,5}
  ch := make(chan int)
  for _, v := range arr {
  
    go func(){
      ch <- v * 2
    }(v)
  }

  for i:=0; i < len(arr); i++{
    fmt.Println(<-ch)
  }
}