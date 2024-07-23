package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	binToDec1("100100")
	binToDec2("100100")
	binToDec3("100100")
	binToDec4("100100")
}

// Самый короткий пример (но без проверки на парсинг)
func binToDec1(bin string) {

	n := 0

	for _, r := range bin { // r - тип рун; Значение символа '0' в ASCII-таблице равно 48. Это связано с тем, что символы в ASCII-таблице хранятся в виде чисел; Hex = 30
		n *= 2
		n += int(r - '0')
	}

	fmt.Println(n)

}

// Использование стандартного пакета strconv (strconv.ParseInt(bin, 2, 64))
func binToDec2(bin string) {

	n, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println("ошибка парсинга")
		return
	}

	fmt.Println(n)

}

// Самый логичный пример (но без проверки на парсинг)
func binToDec3(bin string) {
	result := 0
	dec := 1

	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			result += dec
		}
		dec *= 2
	}
	fmt.Println(result)

}

// Самое плохое решение
func binToDec4(bin string) {

	var result int
	var runesArr = []rune(bin) // [49 48 48 49 48 48]

	bb := len(bin) - 1 // 5

	for _, v := range runesArr {

		vdig := fmt.Sprintf("%c", v) // Конвертируется руна v в строку vdig, содержащую символ, представленный этой руной.

		digital, _ := strconv.Atoi(vdig) // Преобразуется строка vdig в целое число digital

		result += digital * int(math.Pow(2, float64(bb))) // К текущему значению result добавляется произведение числа digital и степени числа 2 в степени bb.

		bb-- // Уменьшается значение переменной bb на 1 для обновления степени числа 2 в следующей итерации.

	}
	fmt.Println(result)
}
