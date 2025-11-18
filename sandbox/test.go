package main

import "fmt"

func main() {
	var users []int = make([]int, 3)
	array := [3]int{1, 2, 3}

	users = array[0:3]

	fmt.Println(users)
}
