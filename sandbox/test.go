package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "eats", meal)
}

func main() {

	tom := person{name: "Tom", age: 24}
	tom.print()
	tom.eat("apples")
}
