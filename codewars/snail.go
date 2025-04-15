/*
Сортировка улитки
Для заданного n x n массива вернуть элементы массива,
упорядоченные от самых внешних элементов к среднему элементу,
по часовой стрелке.
*/
package main

import "fmt"

func SnailTwo(snailMatrix [][]int) []int {
	width, height := len(snailMatrix[0]), len(snailMatrix) //ширина высота матрицы
	edges := [2]int{width + 1, height}

	stepCount := width * height //кол-во шагов всей матрицы
	step := 0                   //текущий уровень, количество шагов на уровне
	dx, dy := 1, 0              //траектория движения
	x, y := -1, 0               //начальная точка
	snailNumbers := make([]int, 0, stepCount)

	movingLine := 0

	for len(snailNumbers) < stepCount {
		
		x += dx
		y += dy
		snailNumbers = append(snailNumbers, snailMatrix[y][x])
		step++
		if step == (2*(width+height) - 4) {
			step = 0
			width -= 2
			height -= 2
		}
	}
	return snailNumbers
}

func snail(snaipMap [][]int) []int {
	var result []int
	lines := len(snaipMap[0])*2 - 1
	moves := len(snaipMap[0])
	dx, dy := 1, 0 //траектория движения
	x, y := -1, 0  //начальная точка
	for ; lines > 0; lines-- {
		for range make([]int, moves) {
			x += dx
			y += dy
			result = append(result, snaipMap[y][x])
		}
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
