package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	name     string
	Lastname string
	Age      int    `json:"age"`
	Phone    string `json:"-"`
	Coments  string `json:",omitempty"`
}

func main() {
	p := Person{
		name:     "John",
		Lastname: "Doe",
		Age:      25,
		Phone:    "123456789",
		Coments:  "This is a comment",
	}

	res, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))
}

// LN AGE COM
