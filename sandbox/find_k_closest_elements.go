package main

import (
	"fmt"
)

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	i, j := findClosestElements(array, 5, 1)
	fmt.Println(array[i], array[j])
}

func findClosestElements(arr []int, k, x int) (int, int) {
	i, j := binarySearch(arr, x)

	for k > 0 {
		k--
		if i > 0 || j < len(arr)-1 && comparison(arr[i], arr[j], x) {
			i--
		} else {
			j++
		}
	}

	return i, j
}

func comparison(a int, b int, x int) bool {
	aMod := mod(a - x)
	bMod := mod(b - x)
	if aMod < bMod {
		return true
	}
	if aMod == bMod && a < b {
		return true
	}
	return false
}

func mod(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func binarySearch(arr []int, x int) (i, j int) {
	i, j = 0, len(arr)-1

	for i+1 < j {
		midIndex := (i + j) >> 1
		if arr[midIndex] >= x {
			j = midIndex
		} else {
			i = midIndex
		}
	}
	return
}
