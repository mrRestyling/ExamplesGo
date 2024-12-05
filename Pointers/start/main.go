package main

import "fmt"

func main() {
	toPrint()
	// toChange()
}

func toPrint() {
	var x int = 10
	var p *int = &x

	fmt.Println("Value of x:", x)    // Выводит значение x
	fmt.Println("Address of x:", &x) // Выводит адрес х
	fmt.Println("Value of p:", p)    // Выводит адрес, хранящийся в p
	fmt.Println("Value at p:", *p)   // Выводит значение по адресу, хранящейся в p
	// разыменование
}

func toChange() {

	x := 40
	fmt.Println("Before:", x)

	changeValue(&x, 50)
	fmt.Println("After:", x)

}

func changeValue(val *int, newValue int) {
	*val = newValue
}
