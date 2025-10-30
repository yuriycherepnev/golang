package main

import "fmt"

var tom struct {
	name string
	age  int
}

func main() {
	tom.name = "Tom"
	fmt.Println(tom) // Tom
}
