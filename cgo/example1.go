package cgo

import "fmt"

/*
	#cgo LDFLAGS: -lm
	#include <stdio.h>
	#include <math.h>
	#include "mylib.h"
	int add(int a, int b) {
		int sum = a + b;
		printf("a: %d, b: %d, sum %d\n", a, b, sum);
		return sum;
	}
*/
import "C"

func CGoExample1() {
	sum := C.add(3, 2)
	fmt.Println(sum)
	fmt.Println(C.sqrt(100))
	fmt.Println(C.multiply(10, 20))
  fmt.Printf("rest result: %d",C.rest(10,3))
}