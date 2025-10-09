package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10}

	result := findClosestElements(array, 3, 11)

	fmt.Println(result)
}

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k

	for left < right {
		mid := left + (right-left)/2

		// Сравниваем расстояния от x до arr[mid] и arr[mid+k]
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
