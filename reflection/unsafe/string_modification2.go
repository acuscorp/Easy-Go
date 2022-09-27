package unsafe

import (
	"fmt"
	"reflect"
	"unsafe"
)

// pointing to the same varable
func InitStringModification2(){
s := ""
	b := []byte("goodbye")
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sHdr.Data = bHdr.Data
	sHdr.Len = bHdr.Len
	fmt.Println(s)

	b[0] = 'x'
	fmt.Println(s)
}