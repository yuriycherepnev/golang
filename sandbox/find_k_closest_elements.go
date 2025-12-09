package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	i := FindClosestElements(array, 5, 1200)
	fmt.Println(i)
}

func FindClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	l, r := findBinarySearch(arr, x)

	for k > 0 {
		k--
		if r == n || l >= 0 && (x-arr[l]) <= (arr[r]-x) {
			l--
		} else {
			r++
		}
	}
	return arr[l+1 : r]
}

func findBinarySearch(arr []int, x int) (int, int) {
	i, j := 0, len(arr)-1
	for i+1 < j {
		midIndex := int(uint(i+j)) >> 1
		if arr[midIndex] > x {
			j = midIndex
		} else {
			i = midIndex
		}
	}
	return i, j
}
