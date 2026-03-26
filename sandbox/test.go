package main

import "fmt"

func main() {
	seen := make(map[int]bool)
	seen[2] = true

	fmt.Println(seen[3])
	fmt.Println(seen[2])

}
