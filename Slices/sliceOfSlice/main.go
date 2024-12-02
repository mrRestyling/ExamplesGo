package main

import "fmt"

func main() {
	Cap()
}

func Cap() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sl := arr[2:5:5]

	fmt.Printf("len: %d, cap: %d\n", len(sl), cap(sl))
	
}
