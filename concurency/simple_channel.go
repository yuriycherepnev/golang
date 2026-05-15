package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int)

	go getNumber(intChan)

	number := <-intChan
	fmt.Println(number)
	fmt.Println(intChan)
}

func getNumber(intChan chan int) {
	intChan <- 5
}
