package main

import "fmt"

func IsEven(num int) bool {
	return num%2 == 0
}

func main() {
	fmt.Println(IsEven(4))  // Ожидается: true
	fmt.Println(IsEven(7))  // Ожидается: false
	fmt.Println(IsEven(0))  // Ожидается: true (0 считается чётным числом)
	fmt.Println(IsEven(-2)) // Ожидается: true
	fmt.Println(IsEven(-3)) // Ожидается: false
}
