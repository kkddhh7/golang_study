package main

import (
	"fmt"
)

var num int

func main() {
	fmt.Println("숫자를 입력하세요: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}
}
