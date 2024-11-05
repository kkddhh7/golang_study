package main

import (
	"fmt"
)

func add(inputChan chan int, outputChan chan int) {
	fnum := <-inputChan
	snum := <-inputChan

	outputChan <- fnum + snum
}

func mul(inputChan chan int, outputChan chan int) {
	fnum := <-inputChan
	snum := <-inputChan

	outputChan <- fnum * snum
}

func main() {
	addinputChan := make(chan int)
	addoutputChan := make(chan int)
	mulinputChan := make(chan int)
	muloutputChan := make(chan int)

	go add(addinputChan, addoutputChan)
	go mul(mulinputChan, muloutputChan)

	num1 := 0
	num2 := 0

	fmt.Printf("첫 번째 정수를 입력하세요 : ")
	fmt.Scan(&num1)
	fmt.Printf("두 번째 정수를 입력하세요 : ")
	fmt.Scan(&num2)

	addinputChan <- num1
	addinputChan <- num2
	mulinputChan <- num1
	mulinputChan <- num2

	fmt.Printf("덧셈 결과는 : %d\n", <-addoutputChan)
	fmt.Printf("곱셈 결과는 : %d\n", <-muloutputChan)

}
