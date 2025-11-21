package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{2, 4, 6, 8, 1000, 1003, 1001, 1002, 2000}

	i := sort.Search(len(array), func(i int) bool { return array[i] >= 999 })

	fmt.Println(i)

}
