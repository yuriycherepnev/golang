package main

import "fmt"

func main() {
	fmt.Println(generateTable(6))
}

func generateTable(size int) [][]int {
	multiplicationTable := make([][]int, 0, size)

	for i := 1; i <= size; i++ {
		tableLine := make([]int, 0, size)
		for j := 1; j <= size; j++ {
			tableLine = append(tableLine, j*i)
		}
		multiplicationTable = append(multiplicationTable, tableLine)
	}

	return multiplicationTable
}
