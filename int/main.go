package main

import (
	"fmt"
	"sort"
)

func main() {

	// maxValue()
	countGO()

}

func maxValue() {

	var i64 int64
	var i32 int32

	i64 = 9223372036854775807
	i32 = 2147483647

	fmt.Println(i64 + 1)
	fmt.Println(i32 + 1)
}

type Person struct {
	Name string
	Age  int
}

func countGO() {

	people := []Person{
		{"John", 25},
		{"Alice", 25},
		{"Bob", 30},
	}

	// Стабильная сортировка по возрасту
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)

	// Нестабильная сортировка по возрасту
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
