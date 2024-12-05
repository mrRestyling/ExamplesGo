package main

import "fmt"

func modifySlice(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}

}

func modSliceNoPointer(arr []int) []int {
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i] * 2
	}

	return res
}

func modifyArray(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func main() {
	a := []int{2, 3, 5, 7, 9}
	fmt.Println("Original array:", a)

	modifySlice(&a)
	fmt.Println("Modified array:", a)

	res := modSliceNoPointer(a)
	fmt.Println("Mod second time", res)

	aa := [5]int{}
	copy(aa[:], a)
	modifyArray(&aa)
	fmt.Println("Mod third time", aa)
}
