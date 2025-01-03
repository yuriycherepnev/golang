/*
Напишите функцию, которая принимает целое число в качестве входных данных и возвращает количество битов,
которые равны единице в двоичном представлении этого числа.
*/
package main

import (
	"fmt"
)

func CountSetBits(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func main() {
	num := 1234
	result := CountSetBits(num)
	fmt.Println(result)
}

/*
num&1 == 1 - выделение младшего бита числа и проверка равняется ли он единице

num >>= 1 - побитовый сдвиг вправо. Бит, находящийся в крайнем правом положении, удаляется.
num >>= 1 эквивалентно делению числа num на 2 с округлением вниз до целого числа.
*/
