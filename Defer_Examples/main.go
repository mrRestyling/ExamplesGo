package main

import "fmt"

// ВНИМАНИЕ НА ВОЗВРАТ ЗНАЧЕНИЯ

func tt() (a int) {
	a = 1

	defer func() {

		a++
	}()
	defer func() {
		a = 5
	}()
	return 3
}

func main() {
	fmt.Println(tt())
}
