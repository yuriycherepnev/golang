package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func getVal(getCh <-chan int) {
	fmt.Println(<-getCh)
	wg.Done()
}

func main() {
	getCh := make(<-chan int)
	wg.Add(1)
	go getVal(getCh)
	wg.Wait()
}
