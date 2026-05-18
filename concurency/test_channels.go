package main

import "fmt"

func main() {
	fmt.Printf("%d +", 1)
}

func createChan(n int) chan int {
	ch := make(chan int)
	go func() {
		ch <- n
	}()
	return ch
}
