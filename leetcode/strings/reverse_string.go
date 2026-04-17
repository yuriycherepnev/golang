package main

import "fmt"

func main() {
	str := "Строка"

	bytes := []byte(str)

	fmt.Println(string(bytes))

}

func reverseString(s []byte) {

	fmt.Println(s)

}
