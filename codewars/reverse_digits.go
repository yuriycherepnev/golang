package main

import "fmt"

func main() {

	count := Digitize(5555)
	fmt.Println(count)
}

func Digitize(number int) (slice []int) {
	for {
		slice = append(slice, number%10)
		number /= 10
		if number == 0 {
			return slice
		}
	}
}

/**
TODO первое решение с вычислением длины числа:

func numberLength(number int) (count int) {
	for number != 0 {
		number = number / 10
		count++
	}
	return count
}

func Digitize(number int) []int {
	println(numberLength(number))
	numberSlice := make([]int, 0, numberLength(number))
	for {
		digit := number % 10
		number = number / 10
		numberSlice = append(numberSlice, digit)
		if number == 0 {
			break
		}
	}
	return numberSlice
}

*/
