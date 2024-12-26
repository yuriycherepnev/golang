package main

import "fmt"

func main() {
	fmt.Println(Arithmetic(5, 2, "add"))      // Output: 7
	fmt.Println(Arithmetic(5, 2, "subtract")) // Output: 3
	fmt.Println(Arithmetic(5, 2, "multiply")) // Output: 10
	fmt.Println(Arithmetic(5, 2, "divide"))   // Output: 2.5
}

func Arithmetic(a int, b int, operator string) float64 {
	result := 0.0
	switch operator {
	case "add":
		return float64(a + b)
	case "subtract":
		result = float64(a - b)
	case "multiply":
		result = float64(a * b)
	case "divide":
		result = float64(a) / float64(b)
	default:
		result = 0
	}
	return result
}
