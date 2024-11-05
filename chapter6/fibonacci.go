package main

// #cgo LDFLAGS: -L. -lfibonacci
// #include <stdio.h>
// int fibonacci(int n);
import "C"
import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("n을 입력하시오: ")
	fmt.Scanf("%d", &n)

	result := C.fibonacci(C.int(n))
	fmt.Printf("피보나치 %d 번째 수는 %d 입니다\n", n, result)
}
