package main

import "fmt"

type Service struct {
	var1 int
	var2 int
	ff   func()
}

func main() {
	var service Service
	service.var1 = 6745
	service.var2 = 2

	service.ff = func() {
		fmt.Println(111)
	}
	fmt.Println(service)
	service.ff()

	colors := [3]int{2: 3}
	fmt.Println(colors) // blue
}
