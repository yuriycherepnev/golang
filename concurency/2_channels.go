package main

import "fmt"

func main() {
	fmt.Println("Start")
	fmt.Println(<-createChan(5))
	fmt.Println("End")
}

func createChan(n int) chan int {
	ch := make(chan int)
	go func() {
		ch <- n
	}()
	return ch
}
