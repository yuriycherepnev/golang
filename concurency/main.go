package main

import (
	"fmt"
	"sync"
)

func goo(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("goo")
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("go func()")
	}()

	wg.Add(1)
	go goo(wg)
	fmt.Println("main()")

	wg.Wait()

}
