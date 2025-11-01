package main

import "fmt"

type person struct { // вложенная структура person
	name string
	age  int
}

type account struct {
	login      string
	password   string
	personInfo person
}

func main() {
	tom := account{
		login:    "tom@hmail.com",
		password: "12345678",
		personInfo: person{
			name: "Tom",
			age:  41,
		},
	}
	fmt.Println(tom)
	fmt.Println("Name: ", tom.personInfo.name) // Name: Tom
}
