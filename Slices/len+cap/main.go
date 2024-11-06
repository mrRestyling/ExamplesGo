package main

import "fmt"

func main() {
	a := make([]int, 5)    // {0,0,0,0,0}
	b := make([]int, 0, 5) // {,,,,}

	fmt.Println(a, b)
}
