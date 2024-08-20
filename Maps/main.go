package main

import (
	"fmt"
	"sort"
)

// создайте функцию getSortedKeys() ниже
// эта функция должна принимать мапу, а возвращать отсортированный слайс ключей этой мапы
// она нужна, чтобы можно было вывести ключи мапы в отсортированном порядке
// так как мапа сама по себе всегда неотсортирована
func getSortedKeys(moves map[int]string) (arr []int) {
	//var arr []int
	for i := range moves {
		arr = append(arr, i)
	}
	sort.Ints(arr)
	return arr
}

func main() {

	// praktikum()

	// praktikumVebinar()

	praktikumVebinarPrimer2()
}

func praktikum() {
	moves := map[int]string{
		1991: "Терминатор 2: Судный день",
		1984: "Терминатор",
		2009: "Терминатор: Да придет спаситель",
		2003: "Терминатор 3: Восстание машин",
		2015: "Терминатор: Генезис",
		2019: "Терминатор: Темные судьбы",
	}

	// delete(moves, 2019)

	if _, ok := moves[2015]; ok {
		fmt.Println("sqklslq")

	}

	// fmt.Println(value)

	sortedMoves := getSortedKeys(moves)

	fmt.Println()
	fmt.Println("Как смотреть франшизу Терминатор:")
	fmt.Println()

	for _, year := range sortedMoves {
		fmt.Println("Год:", year, "Фильм:", moves[year])
	}
}

func praktikumVebinar() {
	moves := map[int]int{
		1991: 1,
		1984: 1,
		2009: 1,
		2003: 1,
		2015: 1,
		2019: 1,
	}

	fmt.Println(moves[1990])

	_, ok := moves[1990]
	if !ok {
		fmt.Println("false")
	}
}

func praktikumVebinarPrimer() {

	arr := []int{1, 2, 3, 1, 1, 23, 4, 5, 64, 7, 48, 97, 3, 4}

	myMap := make(map[int]int)

	for _, n := range arr {
		myMap[n]++
	}

	var arr2 []int

	for k, n2 := range myMap {
		if n2 > 1 {
			arr2 = append(arr2, k)
		}
	}
	fmt.Println(arr2)

}

func praktikumVebinarPrimer2() {
	m := make(map[string]interface{})

	m["int"] = 1
	m["string"] = "hello"
	m["slice"] = []int{1, 2, 3}
	m["map"] = map[string]int{"a": 1, "b": 2}
	m["chan"] = make(chan int)
	m["func"] = func() {}
	m["interface"] = interface{}(1)
}
