package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func main() {
	start := time.Now()
	const number = 20
	fibNumbers := make([]int, 0, 20)
	for i := 0; i <= number; i++ {
		fibNumbers = append(fibNumbers, fibonacci(i))
	}
	fmt.Println(fibNumbers)
	duration := time.Since(start).Seconds()
	fmt.Println(duration)
}
