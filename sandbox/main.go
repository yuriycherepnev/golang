package main

import "fmt"

func main() {
	var a0 int = 0
	var a1 int = 1
	var a2 int = 2
	var a3 int = 3
	var a4 int = 90
	var a5 int = 45

	var b int = a4 ^ a5 // 110 & 010 = 10  - 2

	fmt.Printf("%b\n", a0)
	fmt.Printf("%b\n", a1)
	fmt.Printf("%b\n", a2)
	fmt.Printf("%b\n", a3)
	fmt.Printf("%b\n", a4)
	fmt.Printf("%b\n", a5)

	fmt.Printf("%b\n", b)
	println(b)
}
