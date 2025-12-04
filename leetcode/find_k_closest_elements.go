package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}

	slice := findClosestElements(array, 5, 4)

	fmt.Println(slice)
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	i := sort.SearchInts(arr, x)

	l, r := i-1, i
	for k > 0 {
		k--
		fmt.Println(l, r)
		fmt.Println(r, " == ", n, " || ", l, " >= ", 0, " && (", x, "-", arr[l], ") <= (", arr[r], "-", x, ")")
		if r == n || l >= 0 && (x-arr[l]) <= (arr[r]-x) {
			l--
		} else {
			r++
		}
		fmt.Println(l, r)
	}
	return arr[l+1 : r]
}
