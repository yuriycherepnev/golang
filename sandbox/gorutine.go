package main

import "fmt"

func main() {

	for i := 1; i < 7; i++ {
		go sum1(i)
	}
	fmt.Scanln() // ждем ввода пользователя
	fmt.Println("The End")
}

func sum1(n int) {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	fmt.Println(n, "-", result)
}
