package main

import "fmt"

// Что выведет программа?

func main() {
	var ptr *struct{}
	var iface interface{}
	iface = ptr
	if iface == nil {
		println("It's nil!")
	}

	fmt.Printf("Value: %v,Type: %T\n", iface, iface)
	fmt.Println(iface)
}
