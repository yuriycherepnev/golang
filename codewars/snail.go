/*
Сортировка улитки
Для заданного n x n массива вернуть элементы массива,
упорядоченные от самых внешних элементов к среднему элементу,
по часовой стрелке.
*/

package main

import "fmt"

/*
func Snail(snailMatrix [][]int) []int {
	width, height := len(snailMatrix[0]), len(snailMatrix)
	stepCount := width * height
	snailNumbers := make([]int, 0, stepCount)
	level, levelStep := 0, 0

	for len(snailNumbers) < stepCount {
		if width <= 0 || height <= 0 {
			break
		}

		x, y := calculateCoordinates(level, levelStep, width, height)
		snailNumbers = append(snailNumbers, snailMatrix[y][x])
		levelStep++

		if levelStep == (2*(width+height) - 4) {
			level++
			levelStep = 0
			width -= 2
			height -= 2
		}
	}

	return snailNumbers
}

func calculateCoordinates(level, step, width, height int) (int, int) {
	x, y := level, level
	maxX, maxY := width+level-1, height+level-1

	switch {
	case step < maxX-level:
		x += step
	case step < (maxX-level)+(maxY-level):
		x = maxX
		y += step - (maxX - level)
	case step < (maxX-level)*2+(maxY-level):
		x = maxX - (step - (maxX - level) - (maxY - level))
		y = maxY
	default:
		y = maxY - (step - (2*(maxX-level) + (maxY - level)))
	}

	return x, y
}
*/

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
		fmt.Println("line:", lines)
		fmt.Println("moves:", moves)
		fmt.Println("dx:", dx, "/ dy:", dy)

		for range make([]int, moves) {
			x += dx //1
			y += dy //0
			result = append(result, snaipMap[y][x])
			fmt.Println("x:", x, "/ y:", y)
		}

		fmt.Println("------------------------------------")
		moves -= 1 & dx
		dx, dy = -dy, dx
	}

	return result
}

func main() {
	//array1 := [][]int{
	//	{1, 2, 3, 4},
	//	{12, 13, 14, 5},
	//	{11, 16, 15, 6},
	//	{10, 9, 8, 7},
	//}

	array2 := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
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

	result1 := snail(array2)
	fmt.Println(result1)

	//result2 := snail(array2)
	//fmt.Println(result2)
	//
	//result3 := snail(array3)
	//fmt.Println(result3)
}
