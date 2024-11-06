package main

import (
	"fmt"
	"strings"
)

func main() {

	var text string = "Egor going to Oleg, take beer and go dance with Alisa"

	fmt.Println(Text(text))

}

func Text(text string) string {

	nameMap := make(map[string]int)
	noNameMap := make(map[string]int)

	textWord := strings.Split(text, " ")

	for _, str := range textWord {
		textWordRune := []rune(str)

		if textWordRune[0] >= 65 && textWordRune[0] <= 90 {
			tt := strings.Trim(string(textWordRune), ",.!?:;")

			nameMap[tt]++

		} else {

			noNameMap[str]++

		}

	}

	var strArr []string
	for tStr, _ := range nameMap {
		strArr = append(strArr, tStr)
	}

	result := strings.Join(strArr, " ")

	return result

}
