package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{2, 4, 6, 8, 9, 1000, 1002, 1004, 2000}

	fmt.Println(array[0:9])
	slice := findClosestElements(array, 5, 10)

	fmt.Println(slice)
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	i := sort.SearchInts(arr, x)

	l, r := i-1, i
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
