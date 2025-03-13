/*
Сортировка улитки
Для заданного n x n массива вернуть элементы массива,
упорядоченные от самых внешних элементов к среднему элементу,
по часовой стрелке.
*/

package main

import "fmt"

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

func main() {
	array1 := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	array2 := [][]int{
		{1, 2, 3, 4, 5},
		{18, 19, 20, 21, 6},
		{17, 28, 29, 22, 7},
		{16, 27, 30, 23, 8},
		{15, 26, 25, 24, 9},
		{14, 13, 12, 11, 10},
	}

	array3 := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}

	array4 := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6},
	}

	result1 := Snail(array1)
	fmt.Println(result1)

	result2 := Snail(array2)
	fmt.Println(result2)

	result3 := Snail(array3)
	fmt.Println(result3)

	result4 := Snail(array4)
	fmt.Println(result4)
}
