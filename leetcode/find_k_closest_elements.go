package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1001, 1002, 2000}

	slice := findClosestElements(array, 5, 1)

	fmt.Println(slice)
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	i := sort.SearchInts(arr, x)

	l, r := i-1, i

	for k > 0 {
		k--
		if r == n || l >= 0 && (x-arr[l]) <= (arr[r]-x) {
			fmt.Println(1)
			l--
		} else {
			fmt.Println(2)
			r++
		}
	}
	return arr[l+1 : r]
}
