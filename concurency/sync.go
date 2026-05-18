// синхронизация горутин с помощью каналов

package main

import "fmt"

func sum(n int, done chan bool) {
	result := 0

	for k := 1; k <= n; k++ {
		result += k
		if k == 1 {
			fmt.Printf("%d", 1)
		} else if k <= n {
			fmt.Printf(" + %d", k)
		}
	}

	fmt.Printf(" = %d\n", result)
	done <- true
}

func main() {
	done := make(chan bool)

	count := 10

	for i := 1; i <= count; i++ {
		go sum(i, done)
	}

	for i := 1; i <= count; i++ {
		<-done
	}
}
