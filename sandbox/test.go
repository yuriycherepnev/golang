package main

import "fmt"

func main() {
	seen := make(map[int]bool)
	seen[2] = true

	fmt.Println(8 << 1)
	fmt.Println(9 << 1)
}
