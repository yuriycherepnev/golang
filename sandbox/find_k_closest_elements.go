package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1002, 1004, 2000}

	a := closestBinarySearch(array, 999)
	fmt.Println(array[a])
}

func closestBinarySearch(arr []int, x int) int {
	i, j := 0, len(arr)-1

	for i+1 < j {
		mid := (i + j) >> 1
		if arr[mid] >= x {
			j = mid
		} else {
			i = mid
		}
	}
	if mod(arr[i]-x) < (arr[j] - x) {
		return i
	}
	return j
}

func mod(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
