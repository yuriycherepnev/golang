package main

import "fmt"

func sum(n uint) uint {
	var result uint = 0
	for n != 0 {
		result += n
		n -= 1
	}
	return result
}

func main() {
	fmt.Println(sum(-4))
}
