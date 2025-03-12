/*
Сортировка улитки
Для заданного n x n массива вернуть элементы массива,
упорядоченные от самых внешних элементов к среднему элементу,
по часовой стрелке.
*/

package main

import "fmt"

func Snail(snailMatrix [][]int) []int {
	height, width := len(snailMatrix), len(snailMatrix[0])
	stepCount := width * height
	numbers := make([]int, 0, stepCount)

	level, levelStep := 0, 0

	for len(numbers) < stepCount {
		if height <= 0 || width <= 0 {
			break
		}

		x, y := calculateCoordinates(level, levelStep, height, width)
		numbers = append(numbers, snailMatrix[y][x])
		levelStep++

		if levelStep == (2*(height+width) - 4) {
			level++
			levelStep = 0
			height -= 2
			width -= 2
		}
	}

	return numbers
}

func calculateCoordinates(level, step, height, width int) (int, int) {
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
		{1, 1, 1, 1},
		{4, 5, 5, 2},
		{4, 6, 6, 2},
		{3, 3, 3, 2},
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
