package main

import (
	"fmt"
)

func main() {
	str := "ен hello"

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
	fmt.Println(len(str))

}
