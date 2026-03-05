package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {

	fmt.Printf("Горутина %d начала выполнение \n", id)
	time.Sleep(2 * time.Second) // имитация долгой работы горутины
	fmt.Printf("Горутина %d завершила выполнение \n", id)
	wg.Done() // сигнализируем, что горутина завершила работу
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины

	// вызываем горутины
	go work(1, &wg)
	go work(2, &wg)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}
