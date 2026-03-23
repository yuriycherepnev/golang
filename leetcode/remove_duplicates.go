// [0,0,1,1,1,2,2,3,3,4]
package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	currentIndex := 0

	for _, value := range nums {
		if value != nums[currentIndex] {
			currentIndex++
			nums[currentIndex] = value
		}
	}

	return currentIndex + 1
}
