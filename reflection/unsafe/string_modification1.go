package unsafe

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)


func InitStringModification1(){
  	s := "hello"
	sHdrData := unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len) // prints 5

	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(uintptr(sHdrData) + uintptr(i)))
		fmt.Print(string(bp))
	}
  runtime.KeepAlive(s)

	fmt.Println()
  
	sHdr.Len = sHdr.Len + 10
	fmt.Println(s)

  	s2 := []int{10, 20, 30}
	sHdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	fmt.Println(sHdr2.Len) // prints 3
	fmt.Println(sHdr2.Cap) // prints 3
	intByteSize := unsafe.Sizeof(s2[0])
	fmt.Println(intByteSize)
	for i := 0; i < sHdr2.Len; i++ {
		intVal := *(*int)(unsafe.Pointer(sHdr2.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
	runtime.KeepAlive(s2)
}
