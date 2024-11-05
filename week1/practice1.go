package main

import (
	"fmt"
)

var (
	fnum     = 0
	snum     = 0
	operator = ""
)

func main() {
	fmt.Println("첫 번째 숫자 입력:")
	fmt.Scan(&fnum)

	fmt.Println("두 번째 숫자 입력:")
	fmt.Scan(&snum)

	fmt.Println("연산자 입력(+, -, *, /):")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println(fnum + snum)
	case "-":
		fmt.Println(fnum - snum)
	case "*":
		fmt.Println(fnum * snum)
	case "/":
		if snum == 0 {
			fmt.Println("error")
		} else {
			fmt.Println(fnum / snum)
		}
	default:
		fmt.Println("error")
	}

}
