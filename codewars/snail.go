/*
алгоритм улитки
Оператор ^= — это оператор побитового исключающего ИЛИ (XOR) в Go.
если значения равны то будет 0, если не равны то будет 1

решение имеет линейную сложность O(n), где n — это количество элементов, которые нужно обработать.
В данном случае, поскольку количество элементов в матрице составляет m * n
(где m — количество строк, а n — количество столбцов), сложность будет O(m * n).

Так как сложность зависит от размера входных данных (в данном случае от числа элементов в матрице),
это считается линейной сложностью относительно количества элементов.

Сложность O(m * n) часто используется для алгоритмов,
которые обрабатывают все элементы двумерных структур данных, таких как матрицы.
*/

package main

import "fmt"

func snail(snailMatrix [][]int) []int {
	result := make([]int, 0, len(snailMatrix)*len(snailMatrix[0]))
	lines := [2]int{len(snailMatrix[0]), len(snailMatrix) - 1}
	dx, dy := 1, 0
	x, y := -1, 0
	toggle := 0

	for len(result) < cap(result) {
		for i := 0; i < lines[toggle]; i++ {
			x += dx
			y += dy
			result = append(result, snailMatrix[y][x])
		}
		lines[toggle]--
		toggle = toggle ^ 1
		dx, dy = -dy, dx
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

	array2 := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}

	array3 := [][]int{
		{1, 2, 3, 4, 5},
		{18, 19, 20, 21, 6},
		{17, 28, 29, 22, 7},
		{16, 27, 30, 23, 8},
		{15, 26, 25, 24, 9},
		{14, 13, 12, 11, 10},
	}

	array4 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	array5 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	result1 := snail(array1)
	fmt.Println(result1)

	result2 := snail(array2)
	fmt.Println(result2)

	result3 := snail(array3)
	fmt.Println(result3)

	result4 := snail(array4)
	fmt.Println(result4)

	result5 := snail(array5)
	fmt.Println(result5)
}
