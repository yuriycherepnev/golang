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
