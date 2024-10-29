package main

import (
	"fmt"
	"math"
)

func main() {

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println(LinearSort(arr, 100))
	fmt.Println(BinarySearch(arr, 100))
	// fmt.Println(JumpSearch(arr, 10))

}

func LinearSort(arr []int, s int) int {
	for i, v := range arr {
		if s == v {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, s int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == s {
			return mid
		} else if arr[mid] < s {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // not found
}

func JumpSearch(arr []int, s int) int {

	low := 0
	jump := int(math.Sqrt(float64(len(arr))))

	for {

		if low >= len(arr) {
			low -= jump
			break
		} else if arr[low] == s {
			return low
		} else if arr[low] > s && low == 0 {
			return -1
		} else if arr[low] > s {
			low -= jump
			break
		}

		low += jump
	}

	for j := low; j < low+jump && j < len(arr); j++ {
		if arr[j] == s {
			return j

		}
	}

	return -1
}
