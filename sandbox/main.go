package main

import "fmt"

type Vehicle interface {
	move()
}

func drive(vehicle Vehicle) {
	vehicle.move()
}

type Car struct{}
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func NewCar() *Car {
	return &Car{}
}

func Add(src []int) {
	src = append(src, 1)
}

func main() {
	arr := []int{1, 2, 3}
	src := make([]int, 1)
	copy(src, arr)

	Add(src)
	fmt.Println(arr) // 1 2 3
	fmt.Println(src) // 1 1

	tesla := NewCar()
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}
