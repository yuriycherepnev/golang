package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	i := findClosestElements(array, 3, 0)
	fmt.Println(i)
}

func findClosestElements(arr []int, k int, x int) []int {
	i := closestBinarySearch(arr, x)
	n := len(arr)

	l, r := i, i+1

	for k > 0 {
		k--
		if r == n || l >= 0 && module(arr[l]-x) <= module(arr[r]-x) {
			l--
		} else {
			r++
		}
	}
	return arr[l+1 : r]
}

func module(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func closestBinarySearch(arr []int, x int) int {
	i, j := 0, len(arr)-1
	for i+1 < j {
		midIndex := int(uint(i+j)) >> 1
		if arr[midIndex] > x {
			j = midIndex
		} else {
			i = midIndex
		}
	}
	if module(arr[i]-x) <= module(arr[j]-x) {
		return i
	} else {
		return j
	}
}
