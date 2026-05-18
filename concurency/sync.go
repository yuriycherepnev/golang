// синхронизация горутин с помощью каналов

package main

import (
	"fmt"
	"sync"
)

func sum(n int, done chan bool, mutex *sync.Mutex) {
	result := 0

	out := ""
	for k := 1; k <= n; k++ {
		result += k
		if k == 1 {
			out = fmt.Sprintf("%d", 1)
		} else if k <= n {
			out += fmt.Sprintf(" + %d", k)
		}
	}

	out += fmt.Sprintf(" = %d\n", result)

	mutex.Lock()
	fmt.Print(out)
	mutex.Unlock()

	done <- true
}

func main() {
	done := make(chan bool)
	mutex := &sync.Mutex{}

	count := 10

	for i := 1; i <= count; i++ {
		go sum(i, done, mutex)
	}

	for i := 1; i <= count; i++ {
		<-done
	}
}
