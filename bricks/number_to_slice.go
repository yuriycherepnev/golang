package main

import "fmt"

func main() {
	number := 56784

	result := numberToSlice(number)
	fmt.Println(result)
}

func numberToSlice(number int) []int {
	if number == 0 {
		return []int{0}
	}
	digitCount := 0

	for temp := number; temp != 0; temp /= 10 {
		digitCount++
	}

	numberSlice := make([]int, digitCount)

	for i := digitCount - 1; i >= 0; i-- {
		remainder := number % 10
		number /= 10
		numberSlice[i] = remainder
	}

	return numberSlice
}
