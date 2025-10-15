package main

import "fmt"

func main() {
	age, _ := add(4, 5, "Tom", "Simpson")
	fmt.Println(age) // 9
}

func add(x, y int, firstName, lastName string) (z int, fullName string) {
	z = x + y
	fullName = firstName + " " + lastName
	return
}

/*
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
	return [...]string{"Pending", "Approved", "Rejected", "Cancelled"}[s]
}

func main() {
	var status Status = Approved
	fmt.Println(status)      // "Approved"
	fmt.Println(int(status)) // 1
}

*/
