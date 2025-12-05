package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	i := DownBinarySearch(array, 11)
	fmt.Println(i)
}

func FindClosestElements(arr []int, k, x int) (int, int) {
	n := len(arr)
	l, r := ClosestBinarySearch(arr, x)

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
