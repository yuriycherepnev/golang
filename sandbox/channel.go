package main

import "fmt"

func main() {

	intCh := make(chan int)

	go receive_number(5, intCh) // вызов горутины
	fmt.Println(<-intCh)        // получение данных из канала
	fmt.Println("The End")
}

func receive_number(n int, ch chan int) {

	// отправка данных в канал
	ch <- n * n // отправляем квадрат числа
}
