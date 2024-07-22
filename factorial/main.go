package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

//
func factorial(num uint) uint { // uint - "unsigned integer" (беззнаковое целое число)

	if num == 0 { // Это базовый случай рекурсии. Факториал 0 равен 1
		return 1
	}

	// Здесь функция вызывает саму себя с аргументом num-1.
	// Это рекурсивный шаг, где функция умножает num на результат вызова factorial(num-1).
	// Рекурсия продолжается до тех пор, пока не дойдет до базового случая
	return num * factorial(num-1)

}
