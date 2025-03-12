package main

import "fmt"

func snail(matrix [][]int) [][]int {
	height := len(matrix)
	width := len(matrix[0])

	level := 1
	stepCount := width * height
	levelSize := ((height + width) * 2) - 4
	levelStep := 1

	for i := 1; i <= stepCount; i++ {
		x, y := calculateCoordinates(level-1, levelStep, height, width)
		fmt.Println(x, y)
		levelStep++
		levelSize--
		if levelSize == 0 && i != stepCount {
			height -= 2
			width -= 2
			levelSize = ((height + width) * 2) - 4
			level++
			levelStep = 1
		}
	}

	return matrix
}

func calculateCoordinates(levelIncrease int, levelStep int, height int, width int) (int, int) {
	x := 1
	y := 1
	sideLength := (width + height) - 1
	if levelStep < sideLength {
		if levelStep > width-1 {
			x = width
			y = y + (levelStep - width)
		} else {
			x = levelStep
		}
	} else {
		y = height
		levelStep = levelStep - sideLength
		if levelStep > width-1 {
			y = y - (levelStep - (width - 1))
		} else {
			x = width - levelStep
		}
	}

	return x + levelIncrease, y + levelIncrease
}

func main() {
	//array1 := [][]int{
	//	{1, 2, 3, 5},
	//	{1, 2, 3, 5},
	//	{1, 2, 3, 5},
	//	{1, 2, 3, 5},
	//}

	array1 := [][]int{
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
		{1, 2, 3, 2, 3},
	}

	//array3 := [][]int{
	//	{1, 2, 3, 2},
	//	{1, 2, 3, 2},
	//	{1, 2, 3, 2},
	//}

	result1 := snail(array1)
	fmt.Println(result1)

	//result2 := snail(array2)
	//fmt.Println(result2)
	//
	//result3 := snail(array3)
	//fmt.Println(result3)
}
