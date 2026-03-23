// Дан массив чисел, содержащий минимум два элемента. Нужно найти максимальное произведение двух элементов в этом массиве.
// Input: nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// Output: 16

package mentor

import "fmt"

func main() {
	var sum = nums([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(sum)
}

func nums(numbers []int) int {
	firstMin, secondMin := 0, 0
	firstMax, secondMax := 0, 0
	for _, value := range numbers {
		if value <= firstMin {
			secondMin = firstMin
			firstMin = value
		}
		if value >= firstMax {
			secondMax = firstMax
			firstMax = value
		}
	}
	minSum := firstMin * secondMin
	maxSum := firstMax * secondMax
	if minSum > maxSum {
		return minSum
	} else {
		return maxSum
	}
}
