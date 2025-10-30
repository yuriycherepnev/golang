package main

import "fmt"

type person struct {
	name []int
	age  int
}

func main() {
	var tom person
	tom.age = 66
	tom.name = append(tom.name, 1)
	fmt.Println(tom)
	fmt.Println(tom.name)
}
