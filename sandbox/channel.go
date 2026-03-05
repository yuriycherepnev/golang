package main

import "fmt"

func main() {
	intCh := make(chan int)

	go func() {
		fmt.Println("Go routine starts")

		intCh <- 5
	}()
	chanData := <-intCh

	fmt.Println(chanData)
	fmt.Println("The End")
}
