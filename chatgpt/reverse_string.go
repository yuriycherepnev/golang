package main

import "fmt"

//простое решение
/*
func ReverseString(input string) (result string) {
	for _, v := range input {
		result = string(v) + result
	}
	return result
}
*/

//эффективное решение
func ReverseString(input string) (result string) {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return result
}

func main() {
	fmt.Println(ReverseString("hello"))
}
