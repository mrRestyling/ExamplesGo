package main

import "fmt"

func main() {
	scanFullName()
}

func scanFullName() {
	var firstName, secondName string

	fmt.Println("Введите Ваше имя:")
	fmt.Scanf("%s\n", &firstName)

	fmt.Println("Введите Вашу фамилию:")
	fmt.Scanf("%s\n", &secondName)

	fmt.Printf("Ура, нас посетил:\n -%s %s-", firstName, secondName)
}
