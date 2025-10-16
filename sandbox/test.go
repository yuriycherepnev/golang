package main

import "fmt"

func main() {
	array := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	mid, _ := binarySearch(array, 100)
	k := 3
	slice := make([]int, 0, k)

	for k > 0 {
		slice = append([]int{}, slice...)
		k--
		
	}

	fmt.Println(array[mid])
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
