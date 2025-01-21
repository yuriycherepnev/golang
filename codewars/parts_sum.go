package main

import "fmt"

func main() {
	var number = []uint64{30, 20, 10, 0}
	fmt.Println(PartsSums(number))
}

func PartsSums(ls []uint64) []uint64 {
	partsSums := make([]uint64, 0, len(ls))
	for i := 0; i < len(ls); i++ {
		var sum uint64 = 0
		for j := i; j < len(ls); j++ {
			sum += ls[j]
		}
		partsSums = append(partsSums, sum)
	}
	return partsSums
}
