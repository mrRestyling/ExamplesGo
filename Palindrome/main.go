package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(235555555532))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	n := strconv.Itoa(x)

	for i, numStr := range n {
		if numStr != rune(n[len(n)-1-i]) {
			return false
		}
	}
	return true
}
