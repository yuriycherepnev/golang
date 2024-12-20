package main

import (
	"fmt"
	"math"
)

/*
Если число n составное, его делители состоят из одного числа < чем кв корень из n
и > чем кв корень из n
*/

func IsPrime(num int) bool {
	// Числа <= 1 не являются простыми
	if num <= 1 {
		return false
	}

	// Проверяем делимость от 2 до sqrt(n)
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false // Если делится на i без остатка, число составное
		}
	}

	return true // Если ни на что не делится, число простое
}

func main() {
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(11))
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(-5))
	fmt.Println(IsPrime(97))
}
