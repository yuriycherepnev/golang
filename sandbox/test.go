package main

import "fmt"

func main() {
	// Создаем map и кладем несколько значений
	myMap := make(map[string]int)
	myMap["apple"] = 5
	myMap["banana"] = 3
	myMap["cherry"] = 8

	array := []int{1, 2, 3, 4}
	number := 944

	fmt.Printf("%d\n", number)
	// %b\n byte
	// %c\n runes
	fmt.Printf("%#v\n", array)

	// %b\n byte
	// %c\n runes
}
