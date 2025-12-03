package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1001, 1002, 2000}
	fmt.Println(array)
}

func comparison(a int, b int, x int) int {
	aMod := mod(a - x)
	bMod := mod(b - x)
	if aMod < bMod {
		return a
	}
	if aMod == bMod && a < b {
		return a
	}
	return b
}

func mod(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
