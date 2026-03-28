package main

import "fmt"

func main() {
	result := test()
	fmt.Println(result)
}

func test() (result int) {
	defer func() {
		result++
	}()
	return 1
}

/*

for 1
defer 1 0
defer 2 1
defer 3 2

for 2
defer 1 0
defer 2 1
defer 3 2

*/
