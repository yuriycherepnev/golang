package main

import "fmt"

func main() {
	dx := 1
	moves := 0

	moves -= 1 & dx
	fmt.Println(1 & dx)
	dx = 0
	moves -= 1 & dx

	fmt.Println(1 & dx)
	fmt.Println(1 & -1)
}
