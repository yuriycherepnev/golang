package main

import "fmt"

type person struct {
	name string
	age  int
}

type account struct {
	login      string
	password   string
	personInfo person
}

func main() {
	tom := person{name: "Tom", age: 22}
	var pAge *int = &tom.age
	*pAge = 35
	fmt.Println(tom.age)
}
