package main

import "fmt"

func main() {
	seen1 := "123123"
	seen2 := "фывячс"
	fmt.Printf("%b\n", byte(1))

	for i := 0; i < len(seen1); i++ {
		fmt.Printf("%b\n", seen1[i])
	}

	for i := 0; i < len(seen2); i++ {
		fmt.Printf("%b\n", seen2[i])
	}
}
