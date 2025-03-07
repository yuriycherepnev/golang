package main

import "fmt"

func snail(matrix [][]int) [][]int {
	height := len(matrix)
	width := len(matrix[0])
	x := 0
	y := 0
	i := 1

	stepCount := len(matrix) * len(matrix[0])

	for i < step {

		if step < width {
			y = 0
			x = i
		}

		i++
	}

	return matrix
}

func coordinateCalc(step int, height int, width int) (int, int) {
	x := 0
	y := 0

	if step < width {
		y = 0
		x = step
	}
	if step > width

	return x, y
}

func main() {

	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result := snail(array)
	fmt.Println(result)
}
