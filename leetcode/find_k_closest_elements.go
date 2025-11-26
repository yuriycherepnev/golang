package main

import (
	"fmt"
	"sort"
)

func main() {
	a := -10
	b := -1

	fmt.Println(a * b) // 1

	array := []int{2, 4, 6, 8, 1000, 1003, 1001, 1002, 2000}

	sort.SearchInts(array, a)

	i := Search(len(array), func(i int) bool {
		return array[i] >= 2002
	})

	fmt.Println(i)
}

func Search(n int, function func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := (i + j) >> 1
		if !function(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func comparison(a, b, x int) bool {
	return module(a-x) < module(b-x) || module(a-x) == module(b-x) && a < b
}

func module(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}
