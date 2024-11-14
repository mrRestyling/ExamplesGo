package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(8))
	fmt.Println(fibonacciARR(8))

}

func fibonacci(num int) int {

	x := 0

	y := 1

	for i := 0; i < num; i++ {

		x, y = y, x+y

	}

	return y - x

}

func fibonacciARR(num int) []int {

	x := 0
	y := 1

	arr := make([]int, 0, num)

	for i := 0; i < num; i++ {

		x, y = y, x+y

		arr = append(arr, y-x)

	}

	return arr
}
