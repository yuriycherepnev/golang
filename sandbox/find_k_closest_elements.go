package main

import (
	"fmt"
	"sort"
)

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

func main() {
	array := []int{2, 4, 6, 8, 10, 1000, 1001, 1002, 2000}
	//mid, _ := binarySearch(array, 100)
	//k := 3
	//addNumber := mid
	//
	//slice := make([]int, 0, k)
	//
	//for k > 0 {
	//	left := 0
	//	right := len(array) - 1
	//
	//	slice = append([]int{}, slice...)
	//	k--
	//}

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
	left := 0
	right := len(arr)

	for left+1 != right {
		mid := (left + right) / 2

		if arr[mid] > x {
			right = mid
		} else {
			left = mid
		}
	}

	return left, right
}

/**
package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10}

	result := findClosestElements(array, 3, 11)
	one := 1
	two := 2

	add(one, two)
	add(one, two, one, two)

	fmt.Println(result)
}

func add(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k

	for left < right {
		mid := left + (right-left)/2

		// Сравниваем расстояния от x до arr[mid] и arr[mid+k]
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
*/
