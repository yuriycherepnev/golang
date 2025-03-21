package main

import "fmt"

func main() {
	dx := 1
	moves := 0
	moves -= 1 & dx
	fmt.Println(moves)
	moves -= 1 & dx
	fmt.Println(moves)
	dx = 0
	moves -= 1 & dx
	moves -= 1 & dx
	fmt.Println(moves)
}
