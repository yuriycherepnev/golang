package main

import "fmt"

func main() {
	integers := []int{4, 3, 2, 4, 6}
	fmt.Println(findException(integers))
}

func findException(integers []int) int {
	oddCount, evenCount := 0, 0
	var lastOdd, lastEven int

	for _, number := range integers {
		if number%2 == 0 {
			oddCount++
			lastOdd = number
		} else {
			evenCount++
			lastEven = number
		}
		if oddCount > 1 && evenCount > 0 {
			return lastEven
		}
		if evenCount > 1 && oddCount > 0 {
			return lastOdd
		}
	}

	if evenCount > oddCount {
		return lastOdd
	}
	return lastEven
}
