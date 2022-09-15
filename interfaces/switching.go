package interfaces

import "fmt"

func whatTypeIsThis(i interface{}){
  //fmt.Printf("%T\n", i)
  // v := i.(type) can be used in the switch 
  switch i.(type) {
    case string:
      fmt.Println("It's a string ", i.(string))
    case uint32:
      fmt.Println("It's an unasigned 32-bit integer ", i.(uint32))
    default:
      fmt.Println("Don't know")
  }
}

func InitWhatIsThis(){
  whatTypeIsThis("hola")
  whatTypeIsThis(uint32(3))
  whatTypeIsThis(int64(3))
}