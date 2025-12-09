package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	i := FindClosestElements(array, 5, 6000)
	fmt.Println(i)
}

func FindClosestElements(arr []int, k int, x int) []int {
	l, r := findBinarySearch(arr, x)
	for k > 0 {
		k--
		if l != 0 {
			l--
		} else {
			r++
		}
	}
	return arr[l : r+1]
}

func module(n int) int {
	if n < 0 {
		return -n
	}
	return n
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
