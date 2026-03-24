package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	rotateArray(nums, 4)

	fmt.Println(nums)
}

func rotateArray(nums []int, k int) {
	length := len(nums)
	k = k % length

	firstPart := nums[:length-k]
	secondPart := nums[length-k:]

	copy(nums, append(firstPart, secondPart...))
}
