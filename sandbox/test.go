package main

import (
	"fmt"
	"math"
)

func main() {
	i := math.MaxInt64 - 100
	j := math.MaxInt64

	dev := (i + j) / 2
	dev2 := int(uint(i+j) >> 1)

	fmt.Println(dev)
	fmt.Println(dev2)
}
