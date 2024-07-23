package main

import (
	"fmt"
	"math"
)

func main() {
	fastExample(23, 4)
	trueExample(23, 4)
}

// Использование функции math.Pow стандартной библиотеке Go, которая возводит число в степень
// math.Pow принимает числа в типе данных float64
// float64 - представляет собой 64-битное число с плавающей запятой двойной точности (примерно 15 десятичных знаков)
func fastExample(num, power int) {

	var result float64

	result = math.Pow(float64(num), float64(power))

	fmt.Println(result)
}

// Воспроизведение цикла для возведения в цикл
func trueExample(num, power int) {

	var result int = 1 // любая цифра в степени 0 = 1

	for i := 0; i < power; i++ {
		result *= num // при одной итерации вернет само число (первая степень)

	}

	fmt.Println(result)
}
