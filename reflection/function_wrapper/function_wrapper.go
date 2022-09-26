package function_wrapper

import(
  "fmt"
  "reflect"
  "time"
  "runtime"
)

func InitWrapperFunc(){
  f1 := MakeTimedFunction(timeMe).(func())
  f1()
  f2 := MakeTimedFunction(timeToo).(func(int)int)
  fmt.Println("result ", f2(3))
}

func timeMe(){
  time.Sleep(1 * time.Second)
}
func timeToo(a int) int {
  time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}

func MakeTimedFunction(f interface{}) interface{} {
  rf := reflect.TypeOf(f)
  if rf.Kind() != reflect.Func {
    panic("Expects a function")
  }
  vf := reflect.ValueOf(f)
  wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
    start := time.Now()
    out := vf.Call(in)
    end := time.Now()
    fmt.Printf("Calling %s took %v\n",runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
    return out
  })

  return wrapperF.Interface()
}