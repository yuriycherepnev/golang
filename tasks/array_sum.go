package main

import "fmt"

func SumArray(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(SumArray(nums))              // Ожидается: 15
	fmt.Println(SumArray([]int{0, 0, 0}))    // Ожидается: 0
	fmt.Println(SumArray([]int{-1, -2, -3})) // Ожидается: -6
	fmt.Println(SumArray([]int{10, 20, 30})) // Ожидается: 60
	fmt.Println(SumArray([]int{}))           // Ожидается: 0
}
