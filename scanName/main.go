package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// scanSIMPLE()
	scanBUFIO()

}

func scanSIMPLE() {
	var firstName, secondName string

	fmt.Println("Введите Ваше имя:")
	fmt.Scanf("%s\n", &firstName)

	fmt.Println("Введите Вашу фамилию:")
	fmt.Scanf("%s\n", &secondName)

	fmt.Printf("Ура, нас посетил:\n -%s %s-", firstName, secondName)
}

func scanBUFIO() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int

	fmt.Fscan(in, &num)

	fmt.Fprintln(out, num)
}
