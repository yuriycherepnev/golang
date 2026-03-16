package main

import "fmt"

func main() {
	x := make([]int, 3, 5)

	x = append(x, 1)
	x = append(x, 2)

	y := append(x, 3)
	z := append(x, 4)

	fmt.Println(y)
	fmt.Println(z)
}
