package main

import "fmt"

func snail(matrix [][]int) [][]int {
	height := len(matrix)
	width := len(matrix[0])

	level := 1
	stepCount := width * height
	levelSize := ((height + width) * 2) - 4
	fmt.Println("level: ", level)
	for i := 1; i <= stepCount; i++ {
		fmt.Println("step: ", i)
		levelSize--
		if levelSize == 0 && i != stepCount {
			height -= 2
			width -= 2
			levelSize = ((height + width) * 2) - 4
			level++
			fmt.Println("level: ", level)
		}
	}

	return matrix
}

func main() {

	array1 := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	array2 := [][]int{
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
	}

	array3 := [][]int{
		{1, 2, 3, 2},
		{1, 2, 3, 2},
		{1, 2, 3, 2},
	}

	result1 := snail(array1)
	fmt.Println(result1)

	result2 := snail(array2)
	fmt.Println(result2)

	result3 := snail(array3)
	fmt.Println(result3)
}
