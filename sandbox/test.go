package main

import (
	"fmt"
	"strconv"
)

func main() {
	bin1 := "0111111101011010011000100011010110010011010111010110001001100101"
	bin2 := "0111111101111111011111110111111101111111011111110111111101111111"

	a, _ := strconv.ParseInt(bin1, 2, 64)
	b, _ := strconv.ParseInt(bin2, 2, 64)

	xor := a ^ b

	fmt.Printf("Число A: %s (%d)\n", bin1, a)
	fmt.Printf("Число B: %s (%d)\n", bin2, b)
	fmt.Printf("A ^ B   : %s (%d)\n", strconv.FormatInt(xor, 2), xor)
}
