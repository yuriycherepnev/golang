package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Divisors(30))
	fmt.Println(Divisors2(30))
}

/*
мое решение
*/

func Divisors(number int) int {
	divisor := 1
	divisorsCount := 0
	for divisor <= number {
		if number%divisor == 0 {
			divisorsCount++
		}
		divisor++
	}
	return divisorsCount
}

/*
оптимизированное решение chatGpt
Вместо проверки всех чисел от 1 до n , достаточно проверять только числа до квадратного корня из n
Если число i является делителем, то n/i тоже является делителем. Это значительно сокращает количество итераций.
*/

func Divisors2(number int) int {
	divisorsCount := 0
	sqrt := int(math.Sqrt(float64(number)))

	for i := 1; i <= sqrt; i++ {
		if number%i == 0 {
			// Если i делитель, то число (number / i) тоже делитель
			divisorsCount += 2
		}
	}

	// Если число является точным квадратом, то мы дважды добавили его корень в подсчет
	if sqrt*sqrt == number {
		divisorsCount--
	}

	return divisorsCount
}
