package main

import "fmt"

func main() {
	nums := []int{2, 1, 4, 1, 2}
	single := singleNumber(nums)

	fmt.Println(single)
}

func singleNumber(nums []int) int {
	result := 0

	for _, num := range nums {
		result = result ^ num
	}

	return result
}
