package main

import "fmt"

//func snail(matrix [][]int) [][]int {
//	height := len(matrix)
//	width := len(matrix[0])
//
//	level := 1
//	stepCount := width * height
//	levelSize := ((height + width) * 2) - 4
//	for i := 1; i <= stepCount; i++ {
//		x, y := calculateCoordinates(level, i, width, height)
//		fmt.Println(x, y)
//
//		levelSize--
//		if levelSize == 0 && i != stepCount {
//			height -= 2
//			width -= 2
//			levelSize = ((height + width) * 2) - 4
//			level++
//		}
//	}
//
//	return matrix
//}
//
//func calculateCoordinates(level int, levelStep int, height int, width int) (int, int) {
//	x := 1
//	y := 1
//	sideLength := (width + height) - 1
//
//	if levelStep < sideLength {
//		if levelStep > width-1 {
//			x = width
//			y = y + (levelStep - width)
//		} else {
//			x = levelStep
//		}
//	} else {
//		y = height
//		x = width
//		levelStep = levelStep - sideLength
//
//		if levelStep > width-1 {
//			y = y - (levelStep - (width - 1))
//		} else {
//			x = x - levelStep
//		}
//	}
//
//	return x, y
//}

func snail(snaipMap [][]int) []int {
	var result []int
	lines := len(snaipMap[0])*2 - 1
	moves := len(snaipMap[0])
	dx, dy := 1, 0
	x, y := -1, 0

	for ; lines > 0; lines-- {
		for range make([]int, moves) {
			x += dx
			y += dy
			result = append(result, snaipMap[y][x])
		}
		moves -= 1 & dx
		dx, dy = -dy, dx
		fmt.Println(moves)
	}

	return result
}

func main() {
	array1 := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	//array2 := [][]int{
	//	{1, 2, 3, 2, 3},
	//	{1, 2, 3, 2, 3},
	//	{1, 2, 3, 2, 3},
	//	{1, 2, 3, 2, 3},
	//	{1, 2, 3, 2, 3},
	//	{1, 2, 3, 2, 3},
	//}
	//
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
