package main

import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)

	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })

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

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1001, 1002, 2000}

	fmt.Println(findClosestElements(array, 5, 999))

}

func comparison(a int, b int, x int) int {
	aMod := mod(a - x)
	bMod := mod(b - x)
	if aMod < bMod {
		return a
	}
	if aMod == bMod && a < b {
		return a
	}
	return b
}

func mod(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func binarySearch(arr []int, x int) (int, int) {
	//i, j := 0, len(arr)-1

	//for i < j {
	//	mid := int(uint())
	//}
	return 5, 6

}
