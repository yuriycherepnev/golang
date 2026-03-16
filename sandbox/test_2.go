package main

import "fmt"

func main() {
	x := make([]int, 3, 6)

	fmt.Println(x)
	a := append(x, 1, 2)
	fmt.Println(x)
	b := append(x, 4, 5)
	fmt.Println(x)

	//y := append(x, 3)
	//z := append(x, 4)
	//a := append(x, 5)

	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(z)
}
