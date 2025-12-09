package main

import (
	"fmt"
)

func main() {

	x := 1200
	l := 500
	r := 1501

	fmt.Println(x-l, x-r)

	fmt.Println(module(l-x) <= module(r-x))
}

func module(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
