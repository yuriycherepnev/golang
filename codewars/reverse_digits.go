package main

import (
	"fmt"
	"unsafe"
)

func main() {
	number := 7788554433
	fmt.Println(unsafe.Sizeof(number))

	//count := Digitize(789643)
	//fmt.Println(count)
}

func numberLength(number int) (count int) {
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func Digitize(number int) []int {
	numberSlice := make([]int, 0, numberLength(number))

	for number > 0 {
		digit := number % 10
		number = number / 10
		numberSlice = append(numberSlice, digit)
	}

	return numberSlice
}
