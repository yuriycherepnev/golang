// [0,0,1,1,1,2,2,3,3,4]
package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4}
	numbers = append(numbers[:1], numbers[2:]...)

	fmt.Println(numbers)
}
