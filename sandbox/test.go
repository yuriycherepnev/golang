package main

import "fmt"

func main() {
	seen := make(map[int]bool)
	seen[2] = true

	nums := []int{9, 10, 10, 9, 6}

	single := singleNumber(nums)

	fmt.Println(single)
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		midResult := result ^ num
		fmt.Println(result, num, midResult)
		result = midResult
	}

	return result
}
