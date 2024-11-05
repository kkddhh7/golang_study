package main

import (
	"fmt"
)

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func findMaxMin(arr []int) (int, int) {
	bnum := arr[0]
	snum := arr[0]

	for _, v := range arr {
		if bnum < v {
			bnum = v
		}

		if snum > v {
			snum = v
		}
	}

	return bnum, snum
}

func main() {
	defer fmt.Println("프로그램이 종료되었습니다.")

	arr := []int{3, 5, 1, 2, 0}

	sum := sumArray(arr)
	max, min := findMaxMin(arr)

	fmt.Println("배열의 총합: ", sum)
	fmt.Println("최대값: ", max)
	fmt.Println("최소값: ", min)

	length := len(arr)

	switch {
	case length < 3:
		fmt.Println("배열의 길이가 짧습니다.")
	case length == 3:
		fmt.Println("배열의 길이가 적당합니다.")
	case length > 3:
		fmt.Println("배열의 길이가 깁니다.")
	}

}
