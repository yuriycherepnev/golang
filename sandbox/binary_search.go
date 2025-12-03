package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1001, 1002, 2000}

	//searchIndex := binarySearch(array, 11)

	fmt.Println(array)
}

//func binarySearch(arr []int, x int) int {
//	i, j := 0, len(arr)-1
//
//	for i < j {
//		midIndex := (i + j) >> 1
//		if arr[midIndex] >= x {
//			j = midIndex
//		} else {
//			i = midIndex + 1
//		}
//	}
//	return i
//}
