package main

// #cgo CFLAGS: -O0
// #include <stdio.h>
//
// int numberIsUgly(int x) {
// 	while (x > 1) {
// 		int y = x;
// 		while (y % 2 == 0)
// 			y /= 2;
// 		while (y % 3 == 0)
// 			y /= 3;
// 		while (y % 5 == 0)
// 			y /= 5;
// 		if (x == y)
// 			return 0;
// 		x = y;
// 	}
// 	return 1;
// }
//
// void getNthUglyNumber(int n) {
// 	int i = 0;
// 	int j = 0;
// 	while (j < n) {
// 		i++;
// 		if (numberIsUgly(i)) {
// 			j++;
// 		}
// 	}
// 	printf("%d\n", i);
// }
import "C"

func main() {
	C.getNthUglyNumber(C.int(1000))
}
