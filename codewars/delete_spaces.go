package main

import "fmt"

func main() {
	NoSpace("a1sss aaa")
}

func NoSpace(str string) {
	runes := make([]rune, 0, len(str))
	for index, char := range str {

		fmt.Println(index)
		fmt.Println(char)
	}
	fmt.Println(string(runes))
}
