//package main
//
//import "fmt"
//
//func main() {
//
//	var f func(int, int) int = func(x, y int) int { return x + y }
//
//	v := func(x, y int) int { return x * y }
//
//	fmt.Println(f(3, 4))
//	fmt.Println(v(3, 4))
//}

package main

import "fmt"

type Status int

const (
	Pending   Status = iota // 0
	Approved                // 1
	Rejected                // 2
	Cancelled               // 3
)

func (s Status) String() string {
	names := []string{
		"Pending",
		"Approved",
		"Rejected",
		"Cancelled",
	}
	return names[s]
}

func processServerData(wrongData [...]string) { // Ошибка!
	// Нельзя принимать [...] в параметрах
}

// ОШИБКА: не скомпилируется

func main() {
	var status Status = 0
	fmt.Println(status)      // "Approved"
	fmt.Println(int(status)) // 1

	var v1 = [...]string{"Pending", "Approved", "Rejected", "Cancelled"}[2]

	fmt.Println(v1) // 1
}
